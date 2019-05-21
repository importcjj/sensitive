package main

import (
	"fmt"
	"github.com/importcjj/sensitive"
)

func keywordFilterSearch(content string) (bool, string) {
	filter := sensitive.New()
	filter.AddWord("hello")
	// filter.LoadWordDict("./conf/dict.txt")
	code, word := filter.FindIn(content)
	fmt.Println(content)
	fmt.Printf("[%s]\n", word)
	fmt.Printf("[%v]\n", code)
	return code, word
}

func main() {
	content := `录制 文字 吃猪肉
sdaddasd
dadasd
dada
教性
`
	keywordFilterSearch(content)
}
