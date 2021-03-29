package c

import (
	"fmt"
	"icode.baidu.com/baidu/game-go-lib/inject"
)

type IC interface {
	CC()
}
type C struct {
	ContentC string
	Ctx      *string
}

func init() {
	inject.Container.BindInterface(new(IC), &C{ContentC: "I am C"})
}

func (c *C) CC() {
	fmt.Println(c.ContentC)
}
