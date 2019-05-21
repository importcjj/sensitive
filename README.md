# Sensitive
[![Build Status](https://travis-ci.org/importcjj/sensitive.svg?branch=master)](https://travis-ci.org/importcjj/sensitive)
[![GoDoc](https://godoc.org/github.com/importcjj/sensitive?status.svg)](https://godoc.org/github.com/importcjj/sensitive)

敏感词查找,验证,过滤和替换 🤓 FindAll, Validate, Filter and Replace words. 

#


Usage:

```go
package main

import (
	"fmt"
	"github.com/importcjj/sensitive"
)

func main() {
	filter := sensitive.New()
	filter.LoadWordDict("path/to/dict")
	filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
	filter.AddWord("长者")

	fmt.Println(filter.Filter("我为长者续一秒")) // 我为续一秒
	fmt.Println(filter.Replace("我为长者续一秒", '*')) // 我为**续一秒
	fmt.Println(filter.FindIn("我为长者续一秒"))      // true, 长者
	fmt.Println(filter.Validate("我为长者续一秒"))      // False, 长者
	fmt.Println(filter.FindAll("我为长者续一秒")) // [长者]

	fmt.Println(filter.FindIn("我为长x者续一秒")) // false
	filter.UpdateNoisePattern(`x`)
	fmt.Println(filter.FindIn("我为长x者续一秒")) // true, 长者
	fmt.Println(filter.Validate("我为长x者续一秒"))      // False, 长者
}
```
