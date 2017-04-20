package sensitive

// import (
// 	"fmt"
// )

// Trie 短语组成的Trie树.
type Trie struct {
	Root *Node
}

// Node Trie树上的一个节点.
type Node struct {
	isRootNode bool
	isPathEnd  bool
	Character  rune
	Children   map[rune]*Node
}

// NewTrie 新建一棵Trie
func NewTrie() *Trie {
	return &Trie{
		Root: NewRootNode(0),
	}
}

// Add 添加若干个词
func (tree *Trie) Add(words ...string) {
	for _, word := range words {
		tree.add(word)
	}
}

func (tree *Trie) add(word string) {
	var node = tree.Root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if next, ok := node.Children[r]; ok {
			node = next
		} else {
			newNode := NewNode(r)
			node.Children[r] = newNode
			node = newNode
		}
		if position == len(runes)-1 {
			node.isPathEnd = true
		}
	}
}

// Replace 词语替换
func (tree *Trie) Replace(text string, character rune) string {
	var node = tree.Root
	var parent = tree.Root
	var runes = []rune(text)
	var wordLength = 0

	for position := 0; position < len(runes); position++ {
		r := runes[position]
		next, ok := node.Children[r]
		parent = node

		if !ok {
			if !node.IsRootNode() {
				if wordLength > 0 {
					if parent.IsPathEnd() {
						for i := position - wordLength; i < position; i++ {
							runes[i] = character
						}
					}
				}
				position -= wordLength
			}
			node = tree.Root
			wordLength = 0
			continue
		}

		if position == len(runes)-1 {
			for i := position - wordLength; i <= position; i++ {
				runes[i] = character
			}
		}

		wordLength++
		node = next
	}
	return string(runes)
}

// Filter 词语去除
func (tree *Trie) Filter(text string) string {
	var node = tree.Root
	var parent = tree.Root
	var runes = []rune(text)
	var wordLength = 0
	var result = make([]rune, 0)

	for position := 0; position < len(runes); position++ {
		r := runes[position]
		next, ok := node.Children[r]
		parent = node
		if !ok {
			if !node.IsRootNode() {
				if wordLength > 0 && !parent.IsPathEnd() {
					position -= wordLength
				} else {
					position--
				}
				node = tree.Root
			} else {
				result = append(result, r)
				node = tree.Root
			}
			wordLength = 0
			continue
		}

		wordLength++
		node = next
	}

	return string(result)
}

// FindIn 判断text中是否含有词库中的词
func (tree *Trie) FindIn(text string) (bool, string) {
	var node = tree.Root
	var parent = tree.Root
	var runes = []rune(text)
	var wordLength int
	var word string
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		next, ok := node.Children[r]
		parent = node
		if !ok {
			if !node.IsRootNode() {
				if wordLength > 0 && parent.IsPathEnd() {
					return true, string(runes[position-wordLength : position])
				}

				node = tree.Root
				position -= wordLength
			} else {
				node = tree.Root
			}
			wordLength = 0
			continue
		}
		if next.IsPathEnd() {
			return true, string(runes[position-wordLength : position+1])
		}

		wordLength++
		node = next
	}
	return false, word
}

// NewNode 新建子节点
func NewNode(character rune) *Node {
	return &Node{
		Character: character,
		Children:  make(map[rune]*Node, 0),
	}
}

// NewRootNode 新建根节点
func NewRootNode(character rune) *Node {
	return &Node{
		isRootNode: true,
		Character:  character,
		Children:   make(map[rune]*Node, 0),
	}
}

// IsLeafNode 判断是否叶子节点
func (node *Node) IsLeafNode() bool {
	return len(node.Children) == 0
}

// IsRootNode 判断是否为根节点
func (node *Node) IsRootNode() bool {
	return node.isRootNode
}

// IsPathEnd 判断是否为某个路径的结束
func (node *Node) IsPathEnd() bool {
	return node.isPathEnd
}
