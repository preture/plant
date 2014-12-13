package controllers

import (
	"darling"
	"fmt"
	"github.com/astaxie/beego/orm"
	"net/http"
	"plant/models"
)

type RegisterController struct {
	darling.Controller
}

func (c *RegisterController) Post() {
	req := c.Context.Request
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("path", req.URL.Path)
	fmt.Println("login:", req.Form["login"])
	fmt.Println("password:", req.Form["password"])
	o := orm.NewOrm()
	user := new(models.User)
	user.Login = req.Form["login"][0]
	user.Password = req.Form["password"][0]
	o.Insert(user)
	c.Context.Response.WriteHeader(http.StatusCreated)
}

func (c *RegisterController) Get() {
	c.Post()
}
