//go:generate wecty generate -c TopView -p main top.html
package main

import (
	"syscall/js"

	"github.com/nobonobo/wecty"
)

type TopView struct {
	wecty.Core
}

func (c *TopView) OnSubmit(_ js.Value) interface{} {
	println("submit!")
	return nil
}
