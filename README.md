# Sensitive

敏感词查找,验证,过滤和替换

FindAll, Validate, Filter and Replace words. 

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
	// Do something
}
```

#### AddWord

添加敏感词

```go
filter.AddWord("垃圾")
```

#### Replace

把词语中的字符替换成指定的字符，这里的字符指的是rune字符，比如`*`就是`'*'`。

```go
filter.Replace("这篇文章真的好垃圾", '*')
// output => 这篇文章真的好**
```

#### Filter

直接移除词语

```go
filter.Filter("这篇文章真的好垃圾啊")
// output => 这篇文章真的好啊
```

#### FindIn

查找并返回第一个敏感词，如果没有则返回`false`

```go
filter.FindIn("这篇文章真的好垃圾")
// output => true, 垃圾
```

#### Validate

验证内容是否ok，如果含有敏感词，则返回`false`和第一个敏感词。

```go
filter.Validate("这篇文章真的好垃圾")
// output => false, 垃圾
```

#### FindAll

查找内容中的全部敏感词，以数组返回。

```go
filter.FindAll("这篇文章真的好垃圾")
// output => [垃圾]
```

#### LoadNetWordDict

加载网络词库。

```go
filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
```

#### UpdateNoisePattern

设置噪音模式，排除噪音字符。

```go
// failed
filter.FindIn("这篇文章真的好垃x圾")      // false
filter.UpdateNoisePattern(`x`)
// success
filter.FindIn("这篇文章真的好垃x圾")      // true, 垃圾
filter.Validate("这篇文章真的好垃x圾")    // False, 垃圾
```