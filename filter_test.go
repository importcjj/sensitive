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

	r := filter.Filter("我是共产党p骚宝马娱乐城")
	if r == "我是p" {
		t.Log("func Filter pass")
	} else {
		t.Errorf("expect:%v, actual:%v", "我是p", r)
	}

	r = filter.Replace("我是共产党p骚宝马娱乐城", 42)
	if r == "我是***p******" {
		t.Log("func Replace pass")
	} else {
		t.Error("Replace failed")
	}

	_, result := filter.FindIn("我是共 产 党p骚 宝马 娱乐城")
	if result[0] == "共产党" && result[1] == "骚" && result[2] == "宝马娱乐城" {
		t.Log("func FindIn pass")
	} else {
		t.Error("FindIn failed")
	}

}

func TestAddWord(t *testing.T) {
	filter := New()
	filter.LoadWordDict("path/to/dict.txt")
	filter.AddWord("长者")

	r := filter.Filter("我为长者续一秒")
	if r == "我为续一秒" {
		t.Log("add word and filter pass")
	} else {
		t.Error("add word failed")
	}
	// 42 即 "*"
	r = filter.Replace("我为长者续一秒", 42)
	if r == "我为**续一秒" {
		t.Log("add word and replace pass")
	} else {
		t.Error("add word and replace failed")
	}
	_, result := filter.FindIn("我为长者续一秒")
	if len(result) == 1 && result[0] == "长者" {
		t.Log("add word and findIn pass")
	} else {
		t.Error("add word and findIn failed")
	}

	_, result = filter.FindIn("我为长 者续一秒")
	if len(result) == 1 && result[0] == "长者" {
		t.Log("add word and findIn pass")
	} else {
		t.Error("add word and findIn failed")
	}
}

func TestAddWord2(t *testing.T) {
	filter := New()

	filter.AddWord("习近平下台")
	fmt.Println(filter.Filter("2习近平下台2"))
	fmt.Println(filter.FindIn("2习近平下台2"))
	fmt.Println(filter.Replace("2习近平下台2", 42))
}
