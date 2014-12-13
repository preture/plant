package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:password@localhost/plant")
	orm.DefaultTimeLoc = time.UTC
	orm.RunCommand()
}
