package models

import (
	"time"
)

type User struct {
	Id          int64  `orm:"auto;pk"`
	Login       string `orm:"size(64);unique;index"` //登录帐号
	Password    string
	Email       string    `orm:"unique;index;null"` //邮箱
	Gender      string    `orm:"size(10);null"`
	Birth       time.Time `orm:"type(datetime);null"`
	GravatarId  string    `orm:"size(64);null"`
	GravatarUrl string    `orm:"size(128);null"`
	Location    string    `orm:"null"` //所在位置

	IsStaff   bool      //是否是员工
	IsSuper   bool      //是否是管理员
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	Url       string    `orm:"size(128)"` //访问绝对路径
}

func (u *User) TableName() string {
	return "auth_user"
}
