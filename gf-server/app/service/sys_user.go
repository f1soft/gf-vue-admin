package service

import (
	"errors"
	"gf-server/app/model/user"
	"gf-server/global"
	"github.com/gogf/gf/frame/g"
	uuid "github.com/satori/go.uuid"
)

// Login 用户登录
func Login(u *user.Entity) (userReturn *user.Entity, err error) {
	err = global.GFVA_DB.Table("sys_user").Where(g.Map{"username = ?": u.Username, "password = ?": u.Username}).Struct(&userReturn)
	if userReturn.CompareHashAndPassword(u.Password) { // 检查密码是否正确
		return userReturn, nil
	}
	return u, errors.New("未知用户")
}

// Register 注册
func Register(u *user.Entity) (userReturn *user.Entity, err error) {
	var userFromDb user.Entity
	if err = global.GFVA_DB.Table("sys_user").Where(&u).Struct(&userFromDb); err != nil { // 查询数据是否在数据库
		return userReturn, errors.New("注册失败")
	}
	if userFromDb.UUID == uuid.Nil.String() { // 判断数据的用户UUID是否存在,不存在则插入数据
		u.UUID = uuid.NewV4().String()
		if userReturn, err = u.EncryptedPassword(); err != nil{
			return userReturn, errors.New("注册失败")
		}
		_, err = global.GFVA_DB.Table("sys_user").Data(userReturn).Insert()
		if err != nil {
			return userReturn, errors.New("用户名已注册")
		}
	}
	return userReturn, nil
}

