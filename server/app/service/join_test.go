package service

import (
	"server/app/model"
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

func TestMenu(t *testing.T) {
	authorityMenus := ([]*model.AuthorityMenu)(nil)
	g.DB().SetDebug(true)
	err := g.DB("default").Table("menus a").RightJoin("authority_menu b", "a.id=b.menu_id").Where(g.Map{"authority_id": "888"}).Structs(&authorityMenus)
	if err != nil {
		panic(err)
	}
	//fmt.Println(r)
}

func TestUserAndUserInfo(t *testing.T) {
	user := (*User)(nil)
	userInfo := (*UserInfo)(nil)
	g.DB().SetDebug(true)
	_ = g.DB("test").Table("user").Where(g.Map{"id": 1}).Struct(&user)
	_ = g.DB("test").Table("user_info").Where(g.Map{"uid": 1}).Struct(&userInfo)
	user.Info = *userInfo
}
