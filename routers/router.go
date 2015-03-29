package routers

import (
	"mapcluster/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/maps", &controllers.MainController{}, "get:Maps")
    beego.Router("/hello", &controllers.MainController{}, "get:Hello")
    
}
