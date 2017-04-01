package sensitive_test

import (
	"fmt"
	"github.com/importcjj/sensitive"
)

func ExampleFilter() {
	filter := sensitive.New()
	filter.LoadWordDict("path/to/dict.txt")
	filter.AddWord("长者")

	fmt.Println(filter.Filter("我为长者续一秒"))
	fmt.Println(filter.Replace("我为长者续一秒", 42))
	fmt.Println(filter.FindIn("我为长者续一秒"))
	fmt.Println(filter.FindIn("我为长 者续一秒"))
	// Output:
	// 我为续一秒
	// 我为**续一秒
	// true 长者
	// true 长者
}
