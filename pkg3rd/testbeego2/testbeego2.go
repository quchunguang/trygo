package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type FooController struct {
	beego.Controller
}

func (this *FooController) Get() {
	this.Ctx.WriteString("Bar")
}

func main() {
	// Simplest usage
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})

	beego.Any("/any", func(ctx *context.Context) {
		ctx.Output.Body([]byte("any"))
	})

	// MVC usage
	beego.Router("/foo", &FooController{})

	beego.Run()
}
