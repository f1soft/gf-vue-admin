package service

import (
	"errors"
	"gf-server/app/model/user"
	"gf-server/global"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func Login(u *user.Entity) (userReturn *user.Entity, err error){
	var (
		r gdb.Record
		userInter user.Entity
	)
	r, err = global.GFVA_DB.Table("sys_user").Where(g.Map{"username = ?": u.Username, "password = ?": u.Username}).One()
	err = gconv.Struct(r, &userInter)
	if userReturn.CompareHashAndPassword(u.Password) { // 检查密码是否正确
		return userReturn, nil
	}
	return u, errors.New("未知用户")
}
