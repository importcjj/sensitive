package sensitive

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	tree := NewTrie()

	tree.Add("习近平", "习大大")
	fmt.Println(tree.Root.Children[0])
	fmt.Println(tree.Replace("你好吗 我支持习大大， 他的名字叫做习近平", 42))
	fmt.Println(tree.Filter("你好吗 我支持习大大， 他的名字叫做习近平"))
}
