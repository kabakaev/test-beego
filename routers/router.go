package routers

import (
	"bhi/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/about", &controllers.AboutController{})
    beego.Router("/stat", &controllers.StatController{})
    beego.Router("/stat.png", &controllers.StatGraphController{})
}
