package routers

import (
	"github.com/astaxie/beego"
	"github.com/longan/litchi/web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
