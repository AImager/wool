package main

import (
	"fmt"
	"icode.baidu.com/baidu/game-go-lib/inject"
	"inject/B"
	"inject/D"
)

type E struct {
	ElementB b.IB `inject:""`
	ElementD d.ID `inject:""`
	ContentE string
	Ctx      *string
}

type Test interface {
	TestTest()
}

func (e *E) EE() {
	e.ElementB.BB()
	e.ElementD.DD()
	fmt.Println(e.ContentE)
	fmt.Println(*e.Ctx)
}

func main() {

	ctx := "test"

	e := &E{ContentE: "I am E", Ctx: &ctx}
	inject.Container.Make(e, false)
	e.EE()
	test()
}

func test() Test {
	return &d.D{}
}
