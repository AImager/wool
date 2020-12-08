# wool

<!-- 🇬🇧 English | 🇨🇳 [中文](README_zh.md) -->

## Install

`go get github.com/AImager/wool`

## Usage

```go
package main

import (
	"fmt"
	"github.com/AImager/wool"
)

type AI interface {
	testA()
}

type A1 struct{}

type A2 struct{}

func (a1 *A1) testA() {
	fmt.Println("a1")
}

func (a2 *A2) testA() {
	fmt.Println("a2")
}

func main() {
	// bind interface with factory method
	aFunc := func(name string) AI {
		switch name {
		case "A1":
			return &A1{}
		case "A2":
			return &A2{}
		default:
			return *new(AI)
		}
	}
	_ = wool.DI().Bind(new(AI), aFunc)

	var ai = (AI)(nil)
	_ = wool.DI().Make(&ai, "A2")

	ai.testA()

	// bind interface with factory
	// golang can't find implement struct now
	//_ = di.DI().Bind(new(AI), A1{})
	//
	//ai1 := *new(AI)
	//_ = di.DI().Make(&ai1)
}

```

## Contributing

PRs accepted.

## License

MIT © AImager