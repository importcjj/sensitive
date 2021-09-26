package main

import (
	"fmt"
	"sensitive"
)

func main() {
	filter := sensitive.New()
	filter.LoadWordDict("../../dict/dict2.txt")
	fmt.Println(filter.ValidateWithWildcard("刘一上三台啊", '*'))
	fmt.Println(filter.ValidateWithWildcard("哈哈哈刘一上三", '*'))
	fmt.Println(filter.ValidateWithWildcard("哈哈哈刘一上三台", '*'))
}
