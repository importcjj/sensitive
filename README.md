# Sensitive

敏感词过滤 [![Build Status](https://travis-ci.org/importcjj/sensitive.svg?branch=master)](https://travis-ci.org/importcjj/sensitive)

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
	filter.LoadWordDict("../dict/dict.txt")
	filter.AddWord("长者")

	fmt.Println(filter.Filter("我为长者续一秒")) // 我为续一秒
	// 42 即 "*"
	fmt.Println(filter.Replace("我为长者续一秒", 42)) // 我为**续一秒
	fmt.Println(filter.FindIn("我为长者续一秒"))      // true, 长者

	fmt.Println(filter.FindIn("我为长|者续一秒")) // false,
	filter.UpdateNoisePattern(`\|`)
	fmt.Println(filter.FindIn("我为长|者续一秒")) // true, 长者
}
```
