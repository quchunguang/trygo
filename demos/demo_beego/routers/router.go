package routers

import (
	"github.com/quchunguang/trygo/demos/demo_beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
