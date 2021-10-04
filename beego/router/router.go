package router

import (
	"beego/controller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/goods", &controller.GoodsController{})
	beego.Router("/cms_:id[0-9].html", &controller.ArticleController{})
	beego.Router("/user/get", &controller.UserController{}, "get:GetUser")
	beego.Router("/user/dele", &controller.UserController{}, "get:DeleUser")
	beego.Router("/user/xml", &controller.UserController{}, "post:Xml")
	beego.Router("/user/put", &controller.UserController{}, "put:Put")

}
