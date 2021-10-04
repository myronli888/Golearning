package controller

import (
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["title"] = "你好beego aaa"
	c.TplName = "goods.tpl"
}
