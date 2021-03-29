package a

import (
	"fmt"
	"icode.baidu.com/baidu/game-go-lib/inject"
)

type IA interface {
	AA()
}
type A struct {
	ContentA string
	Ctx      *string
}

func init() {
	inject.Container.BindInterface(new(IA), &A{ContentA: "I am A"})
}

func (a *A) AA() {
	fmt.Println(a.ContentA)
}
