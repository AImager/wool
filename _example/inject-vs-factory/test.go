package main

import (
	"fmt"
)

func main() {
	a := factoryA("A1", "E2")
	a.Print()
}

type interfaceA interface {
	Print()
}

type A1 struct {
	fieldE interfaceE
}

func (a A1) Print() {
	fmt.Println("hello a1")
	a.fieldE.Print()
}

type A2 struct {
	fieldE interfaceE
}

func (a A2) Print() {
	fmt.Println("hello a2")
	a.fieldE.Print()
}

func factoryA(nameA, nameE string) (iA interfaceA) {
	switch nameA {
	case "A1":
		iA = A1{factoryE(nameE)}
	case "A2":
		iA = A2{factoryE(nameE)}
	}
	return
}

type interfaceE interface {
	Print()
}

type E1 struct{}

func (e E1) Print() {
	fmt.Println("hello e1")
}

type E2 struct{}

func (e E2) Print() {
	fmt.Println("hello e2")
}

func factoryE(nameE string) (iE interfaceE) {
	switch nameE {
	case "E1":
		iE = E1{}
	case "E2":
		iE = E2{}
	}
	return
}
