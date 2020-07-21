package service

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/frame/g"
)

type User struct {
	Id       int      `orm:"id" json:"id"`
	Username string   `orm:"username" json:"username"`
	Info     UserInfo `json:"user_info"`
}

type UserInfo struct {
	Uid   int    `orm:"uid" json:"-"`
	Email string `orm:"email" json:"email"`
}

func TestUser(t *testing.T) {
	g.DB().SetDebug(true)
	r, err := g.DB("test").Table("user u").RightJoin("user_info ui", "u.id=ui.uid").Fields("u.id, u.username, ui.email").One()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

func TestUserAndUserInfo(t *testing.T) {
	user := (*User)(nil)
	userInfo := (*UserInfo)(nil)
	g.DB().SetDebug(true)
	_ = g.DB("test").Table("user").Where(g.Map{"id": 1}).Struct(&user)
	_ = g.DB("test").Table("user_info").Where(g.Map{"uid": 1}).Struct(&userInfo)
	user.Info = *userInfo
}
