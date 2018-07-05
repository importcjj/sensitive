package main

import (
	"fmt"
	"github.com/importcjj/sensitive"
)

func main() {
	filter := sensitive.New()
	filter.LoadWordDict("../dict/dict.txt")
	filter.AddWord("一个东西")
	filter.AddWord("一些")
	filter.AddWord("个")
	filter.AddWord("有一")
	filter.AddWord("一个")

	fmt.Println(filter.Replace("hello", 42))
	fmt.Println(filter.Replace("骚", 42))

	fmt.Println(filter.Replace("一", 42))
	fmt.Println(filter.Replace("一个", 42))
	fmt.Println(filter.Replace("一个东", 42))
	fmt.Println(filter.Replace("一个东西", 42))
	fmt.Println(filter.Replace("一个东西啊", 42))
	fmt.Println(filter.Replace("有一个东西啊", 42))
	fmt.Println(filter.Replace("有一个东啊", 42))
	fmt.Println(filter.Replace("有一个啊", 42))
	fmt.Println(filter.Replace("有一个", 42))
	fmt.Println(filter.Replace("有一", 42))

	fmt.Println(filter.Validate("一"))
	fmt.Println(filter.Validate("一个"))
	fmt.Println(filter.Validate("一个东"))
	fmt.Println(filter.Validate("一个东西"))
	fmt.Println(filter.Validate("一个东西啊"))
	fmt.Println(filter.Validate("有一个东西啊"))
	fmt.Println(filter.Validate("有一个东啊"))
	fmt.Println(filter.Validate("有一个啊"))
	fmt.Println(filter.Validate("有一个"))
	fmt.Println(filter.Validate("有一"))

	fmt.Println("一", filter.Filter("一"))
	fmt.Println("一个", filter.Filter("一个"))
	fmt.Println("一个东", filter.Filter("一个东"))
	fmt.Println("一个东西", filter.Filter("一个东西"))
	fmt.Println("一个东西啊", filter.Filter("一个东西啊"))
	fmt.Println("有一个东西啊", filter.Filter("有一个东西啊"))
	fmt.Println("有一个东啊", filter.Filter("有一个东啊"))
	fmt.Println("有一个啊", filter.Filter("有一个啊"))
	fmt.Println("有一个", filter.Filter("有一个"))
	fmt.Println("有一", filter.Filter("有一"))

}
