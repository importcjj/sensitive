package sensitive

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	filter := New()
	if err := filter.LoadWordDict("dict/dict.txt"); err != nil {
		t.Error(err)
	}

	fmt.Println(filter.Filter("我是共产党p骚宝马娱乐城"))
	fmt.Println(filter.Replace("我是共产党p骚宝马娱乐城", 42))
	fmt.Println(filter.FindIn("我是共 产 党p骚 宝马 娱乐城"))
}

func TestAddWord(t *testing.T) {
	filter := New()
	filter.LoadWordDict("path/to/dict.txt")
	filter.AddWord("长者")

	fmt.Println(filter.Filter("我为长者续一秒"))
	// 42 即 "*"
	fmt.Println(filter.Replace("我为长者续一秒", 42))
	fmt.Println(filter.FindIn("我为长者续一秒"))
	fmt.Println(filter.FindIn("我为长 者续一秒"))
}

func TestAddWord2(t *testing.T) {
	filter := New()

	filter.AddWord("习近平下台")
	fmt.Println(filter.Filter("2习近平下台2"))
	fmt.Println(filter.FindIn("2习近平下台2"))
	fmt.Println(filter.Replace("2习近平下台2", 42))
}
