package b

import (
	"fmt"
	"icode.baidu.com/baidu/game-go-lib/inject"
	"inject/A"
)

type IB interface {
	BB()
}
type B struct {
	ElementA a.IA `inject:""`
	ContentB string
	Ctx      *string
}

func (b *B) BB() {
	b.ElementA.AA()
	fmt.Println(b.ContentB)
}

type B1 struct {
	ElementA  a.IA `inject:""`
	ContentB1 string
	Ctx       *string
}

func init() {

	inject.Container.BindInterface(new(IB), &B1{ContentB1: "I am B1"})
}

func (b *B1) BB() {
	b.ElementA.AA()
	fmt.Println(b.ContentB1)
}
