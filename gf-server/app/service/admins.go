package service

import (
	"errors"
	"gf-server/app/api/request"
	"gf-server/app/model/admins"
	"gf-server/library/utils"

	"github.com/gogf/gf/frame/g"
)

// ChangePassword Change administrator password
// ChangePassword 修改管理员密码
func ChangePassword(change *request.ChangePasswordRequest) (err error) {
	var admin *admins.Entity
	if admin, err = admins.FindOne(g.Map{"uuid": change.Uuid}); err != nil {
		return errors.New("用户不存在, 修改密码失败")
	}
	if utils.CompareHashAndPassword(admin.Password, change.Password) {
		if admin.Password, err = utils.EncryptedPassword(change.NewPassword); err != nil {
			if _, err = admins.Save(admin); err != nil {
				return errors.New("修改密码失败")
			}
			return nil
		}
	}
	return errors.New("旧密码输入有误")
}
