package service

import (
	"errors"
	"gf-server/app/model/admins"

	"github.com/gogf/gf/frame/g"
	uuid "github.com/satori/go.uuid"
)

// Login 用户登录
func Login(u *admins.Entity) (userReturn *admins.Entity, err error) {
	admin := admins.GetOne(g.Map{"username = ?": u.Username})
	if admin.CompareHashAndPassword(u.Password) { // 检查密码是否正确
		return userReturn, nil
	}
	return u, errors.New("未知用户")
}

// Register 注册
func Register(u *admins.Entity) (err error) {
	if !admins.RecordNotFound(g.Map{"username": u.Username}) {
		return errors.New("用户已存在,注册失败")
	}
	u.UUID = uuid.NewV4().String()
	if err = u.EncryptedPassword(); err != nil { // 哈希加密
		return errors.New("注册失败")
	}
	if _, err = admins.Insert(u); err != nil {
		return errors.New("注册失败")
	}
	return nil
}
