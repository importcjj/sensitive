# Sensitive

敏感词查找,验证,过滤和替换 FindAll, Validate, Filter and Replace words. 

[![Build Status](https://travis-ci.org/importcjj/sensitive.svg?branch=master)](https://travis-ci.org/importcjj/sensitive) [![GoDoc](https://godoc.org/github.com/importcjj/sensitive?status.svg)](https://godoc.org/github.com/importcjj/sensitive)


新增分支[Aho-Corasick](https://github.com/importcjj/sensitive/tree/Aho-Corasick)以支持AC自动机

#### 用法

```go
package main

import (
	"fmt"
	"github.com/importcjj/sensitive"
)

func main() {
	filter := sensitive.New()
	filter.LoadWordDict("path/to/dict")
	filter.AddWord("垃圾")

	filter.Filter("这篇文章真的好垃圾")       // 这篇文章真的好
	filter.Replace("这篇文章真的好垃圾", '*') // 这篇文章真的好**
	filter.FindIn("这篇文章真的好垃圾")       // true, 垃圾
	filter.Validate("这篇文章真的好垃圾")     // False, 垃圾
	filter.FindAll("这篇文章真的好垃圾")      // [垃圾]
}
```

#### 加载网络词库

```go
filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
```

#### 排除干扰

```go
// failed
filter.FindIn("这篇文章真的好垃x圾")      // false
filter.UpdateNoisePattern(`x`)
// success
filter.FindIn("这篇文章真的好垃x圾")      // true, 垃圾
filter.Validate("这篇文章真的好垃x圾")    // False, 垃圾
```