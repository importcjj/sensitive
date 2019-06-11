package sensitive

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	tree := NewTrie()

	tree.Add("习近平", "习大大")
	tree.BuildFailureLinks()
	fmt.Println(tree.Root.Children[0])
	fmt.Println(tree.Replace("你好吗 我支持习大大， 他的名字叫做习近平", '*'))
	fmt.Println(tree.Filter("你好吗 我支持习大大， 他的名字叫做习近平"))
}

func TestTrieTreeBFS(t *testing.T) {
	tree := NewTrie()
	tree.Add("习近平", "习大大", "共产党好")
	ch := tree.bfs()
	expect := []struct {
		Character string
		Depth     int
	}{
		{"习", 1}, {"共", 1},
		{"近", 2}, {"大", 2}, {"产", 2},
		{"平", 3}, {"大", 3}, {"党", 3},
		{"好", 4},
	}
	i := 0
	for n := range ch {
		if string(n.Character) != expect[i].Character {
			t.Errorf("Expect [%s], got [%s]", expect[i].Character, string(n.Character))
		}
		if n.depth != expect[i].Depth {
			t.Errorf("Expect [%d], got [%d]", expect[i].Depth, n.depth)
		}
		i++
	}
}

func TestTrieTreeBuildFailureLinks(t *testing.T) {
	tree := NewTrie()
	tree.Add("he", "his", "she", "hers")
	tree.BuildFailureLinks()
}
