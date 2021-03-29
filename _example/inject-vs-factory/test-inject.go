package main

import (
	"fmt"
	"icode.baidu.com/baidu/game-go-lib/inject"
)

type IA interface {
	AA()
}
type A struct {
	contentA string
	Ctx      string
}

func (a *A) AA() {
	fmt.Println(a.contentA)
}

type IB interface {
	BB()
}
type B struct {
	ElementA IA `inject:""`
	contentB string
	Ctx      string
}

func (b *B) BB() {
	b.ElementA.AA()
	fmt.Println(b.contentB)
}

type B1 struct {
	ElementA  IA `inject:""`
	contentB1 string
	Ctx       string
}

func (b *B1) BB() {
	b.ElementA.AA()
	fmt.Println(b.contentB1)
}

type IC interface {
	CC()
}
type C struct {
	contentC string
	Ctx      string
}

func (c *C) CC() {
	fmt.Println(c.contentC)
}

type ID interface {
	DD()
}
type D struct {
	ElementC IC `inject:""`
	contentD string
	Ctx      string
}

func (d *D) DD() {
	d.ElementC.CC()
	fmt.Println(d.contentD)
}

type E struct {
	ElementB IB `inject:""`
	ElementD ID `inject:""`
	contentE string
	Ctx      string
}

func (e *E) EE() {
	e.ElementB.BB()
	e.ElementD.DD()
	fmt.Println(e.contentE)
}
func main() {
	inject.Container.BindInterface(new(IB), &B1{contentB1: "I am B1"})
	inject.Container.BindInterface(new(ID), &D{contentD: "I am D"})
	inject.Container.BindInterface(new(IA), &A{contentA: "I am A"})
	inject.Container.BindInterface(new(IC), &C{contentC: "I am C"})
	e := &E{contentE: "I am E"}
	inject.Container.Make(e, false)
	e.EE()
}
