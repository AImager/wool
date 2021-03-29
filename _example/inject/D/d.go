package d

import (
	"fmt"
	"icode.baidu.com/baidu/game-go-lib/inject"
	"inject/C"
)

type ID interface {
	DD()
}

type D struct {
	ElementC c.IC `inject:""`
	ContentD string
	Ctx      *string
}

func init() {
	inject.Container.BindInterface(new(ID), &D{ContentD: "I am D"})
	// inject.Container.BindFunction(new(ID), func(ctx *string) ID {
	// 	return &D{
	// 		ContentD: "I am D",
	// 		Ctx:      ctx,
	// 		// ElementC: inject.Container.Make(new(IC), ctx),
	// 	}
	// })
}

func (d *D) DD() {
	d.ElementC.CC()
	fmt.Println(d.ContentD)
}

func (d *D) TestTest() {
	d.ElementC.CC()
	fmt.Println(d.ContentD)
}
