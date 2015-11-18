package routers

import (
	"github.com/quchunguang/trygo/standalone/testbeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
