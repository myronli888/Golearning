package controller

import (
	"encoding/xml"
	"fmt"

	"github.com/astaxie/beego"
)

type User struct {
	Username string `form:username xml:username`
	Password string `form:password xml:password`
}

type UserController struct {
	beego.Controller
}

func (c *UserController) GetUser() {
	u := User{
		Username: "lisi",
		Password: "135876",
	}
	c.Data["json"] = u
	c.ServeJSON()
}

func (c *UserController) DeleUser() {
	u := User{
		Username: "张三",
		Password: "8888",
	}
	c.Data["json"] = u
	c.ServeJSON()
}

func (c *UserController) Xml() {
	u2 := User{}
	var err error
	if e := xml.Unmarshal(c.Ctx.Input.RequestBody, &u2); e != nil {
		c.Data["Json"] = err.Error()
		c.ServeJSON()
	} else {
		fmt.Printf("%#v", u2)
		c.Data["Json"] = u2
		c.ServeJSON()
	}
}

func (c *UserController) Put() {
	u2 := User{}
	if err := c.ParseForm(&u2); err != nil {
		c.Ctx.WriteString("获取数据失败")
		return
	}
	fmt.Printf("%#v", u2)
	c.Ctx.WriteString("执行修改操作")
}
