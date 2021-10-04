package controller

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("cms 接口" + id)
}
