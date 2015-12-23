React Diff binary tree in Golang
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/react-diff/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/react-diff?status.svg)](https://godoc.org/github.com/kkdai/react-diff)  [![Build Status](https://travis-ci.org/kkdai/react-diff.svg?branch=master)](https://travis-ci.org/kkdai/react-diff)


What is React Diff
---------------


How it works
---------------



Install
---------------
`go get github.com/kkdai/react-diff`


Usage
---------------

```go

package main

import (
	. "github.com/kkdai/react-diff"
)

func main() {
	nT := NewReactDiffTree(20)
	nT.InsertNote("a", 1)
	nT.InsertNote("b", 2)
	nT.InsertNote("c", 3)
	nT.InsertNote("d", 4)
	nT.InsertNote("f", 6)
	nT.InsertNote("e", 8)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNote("a", 1)
	nT2.InsertNote("b", 2)
	nT2.InsertNote("c", 3)
	nT2.InsertNote("d", 5)
	nT2.InsertNote("h", 7)
	nT2.InsertNote("e", 10)

	nT.DiffTree(nT2, INSERT_MARKUP)
	nT.DisplayGraphvizTree()
}
```

![](images/ex1.png)

Inspired
---------------

- [React 源码剖析系列 － 不可思议的 react diff](http://zhuanlan.zhihu.com/purerender/20346379)
- [React -Get Started](http://facebook.github.io/react/docs/getting-started.html)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

