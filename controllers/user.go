package controllers

import (
	"bytes"
	"crypto/md5"
	"darling"
	"encoding/json"
	"fmt"
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
			data, _ := json.Marshal(user)
			c.Context.Response.Header().Add("Content-Type", "application/json; charset=utf-8")
			h := md5.New()
			io.WriteString(h, string(data))
			buffer := bytes.NewBuffer(nil)
			fmt.Fprintf(buffer, "%x\n", h.Sum(nil))
			c.Context.Response.Header().Add("ETag", buffer.String())
			io.WriteString(c.Context.Response, string(data))
		} else {
			c.Context.Response.WriteHeader(http.StatusNotFound)
		}
	}
}

func (c *UserController) Put() {
	id, err := strconv.ParseInt(c.Context.Params[0], 10, 0)
	if err == nil {
		req := c.Context.Request
		req.ParseForm()

		o := orm.NewOrm()
		user := models.User{Id: id}
		if o.Read(&user) == nil {
			user.Gender = req.Form["gender"][0]
			if _, err := o.Update(&user); err == nil {
				user := models.User{Id: id}
				o.Read(&user)
				data, _ := json.Marshal(user)
				c.Context.Response.Header().Add("Content-Type", "application/json; charset=utf-8")
				h := md5.New()
				io.WriteString(h, string(data))
				buffer := bytes.NewBuffer(nil)
				fmt.Fprintf(buffer, "%x\n", h.Sum(nil))
				c.Context.Response.Header().Add("ETag", buffer.String())
				io.WriteString(c.Context.Response, string(data))
			}
		}
	}
}
