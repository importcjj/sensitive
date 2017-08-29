package main

import (
	"fmt"
	"github.com/importcjj/sensitive"
)

func main() {
	filter := sensitive.New()
	filter.AddWord("一个东西")

	fmt.Println(filter.Replace("一个", 42))
}
