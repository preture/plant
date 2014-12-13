package controllers

import (
	"darling"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"io"
	"net/http"
	"plant/models"
)

type RegisterController struct {
	darling.Controller
}

func (c *RegisterController) Post() {
	req := c.Context.Request
	req.ParseForm()
	o := orm.NewOrm()
	user := new(models.User)
	user.Login = req.Form["login"][0]
	user.Password = req.Form["password"][0]
	id, err := o.Insert(user)
	if err == nil {
		createdUser := models.User{Id: id}
		o.Read(&createdUser)
		c.Context.Response.WriteHeader(http.StatusCreated)
		data, _ := json.Marshal(createdUser)
		io.WriteString(c.Context.Response, string(data))
	} else {
		c.Context.Response.WriteHeader(http.StatusForbidden)
	}
}
