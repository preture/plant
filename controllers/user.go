package controllers

import (
	"darling"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"io"
	"net/http"
	"plant/models"
	"strconv"
)

type UserController struct {
	darling.Controller
}

func (c *UserController) Get() {
	id, err := strconv.ParseInt(c.Context.Params[0], 10, 0)
	if err == nil {
		o := orm.NewOrm()
		user := models.User{Id: id}
		err = o.Read(&user)
		if err == nil {
			c.Context.Response.Header().Add("Content-Type", "application/json; charset=utf-8")
			data, _ := json.Marshal(user)
			io.WriteString(c.Context.Response, string(data))
		} else {
			c.Context.Response.WriteHeader(http.StatusNotFound)
		}
	}
}
