package routers

import (
	"github.com/astaxie/beego"
	"github.com/longan/litchi/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
