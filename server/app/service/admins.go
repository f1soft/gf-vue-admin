package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model/admins"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"
)

// ChangePassword Change administrator password
// ChangePassword 修改管理员密码
func ChangePassword(change *request.ChangePassword) (err error) {
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

// UploadHeaderImg User uploads profile picture
// UploadHeaderImg 用户上传头像
func UploadHeaderImg(userUuid string, filePath string) (adminInfo *admins.Entity, err error) {
	if _, err := admins.Update(g.Map{"header_img": filePath}, g.Map{"uuid": userUuid}); err != nil {
		return adminInfo, errors.New("")
	}
	adminInfo, err = admins.FindOne(g.Map{"uuid": userUuid})
	return adminInfo, err
}

// GetAdminInfoList Paging gets the list of users
// GetAdminInfoList 分页获取用户列表
func GetAdminInfoList(info *request.PageInfo) (list interface{}, total int, err error) {
	var adminList *admins.Entity
	db := g.DB("").Table("admins").Safe()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Scan(&adminList)
	return adminList, total, err
}

// SetUserAuthority Set user permissions
// SetUserAuthority 设置用户权限
func SetUserAuthority(set *request.SetAdminAuthority) (err error) {
	_, err = admins.Update(g.Map{"authority_id": set.Uuid}, g.Map{"uuid": set.Uuid})
	return err
}

// DeleteAdmin Delete administrator
// DeleteAdmin 删除管理员
func DeleteAdmin(delete *request.DeleteAdmin) (err error) {
	_, err = admins.Delete(g.Map{"id": delete.Id})
	return err
}
