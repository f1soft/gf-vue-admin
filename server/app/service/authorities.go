package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model"
	"server/app/model/admins"
	"server/app/model/authorities"
	"server/app/model/authority_menu"
	"server/app/model/authority_resources"
	"server/library/global"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"
)

// CreateAuthority Create a role
// CreateAuthority 创建一个角色
func CreateAuthority(create *request.CreateAuthority) (authority *authorities.Entity, err error) {
	insert := &authorities.Entity{
		AuthorityId:   create.AuthorityId,
		AuthorityName: create.AuthorityName,
		ParentId:      create.ParentId,
	}
	if !authorities.RecordNotFound(g.Map{"authority_id": insert.AuthorityId}) {
		return insert, errors.New("存在相同角色id")
	}
	if _, err = authorities.Insert(insert); err != nil {
		return insert, errors.New("创建角色失败")
	}
	return insert, nil
}

// CopyAuthority Copy a character
// CopyAuthority 复制一个角色
func CopyAuthority(copyInfo *request.AuthorityCopy) (authority *authorities.Authorities, err error) {
	var (
		menuList []*model.AuthorityMenu
		baseMenu []model.BaseMenu
	)
	if !authorities.RecordNotFound(g.Map{"authority_id": copyInfo.Authority.AuthorityId}) {
		return authority, errors.New("存在相同角色id")
	}
	info := &request.AuthorityIdInfo{AuthorityId: copyInfo.OldAuthorityId}
	menuList, err = GetMenuAuthority(info)
	for _, v := range menuList {
		v.BaseMenu.Id = gconv.Uint(v.MenuId)
		baseMenu = append(baseMenu, v.BaseMenu)
	}
	copyInfo.Authority.BaseMenu = baseMenu

	//var menusReturn []model.AuthorityMenus
	//authority = &authorities.Authorities{
	//	AuthorityId:   copyInfo.Authority.AuthorityId,
	//	AuthorityName: copyInfo.Authority.AuthorityName,
	//	ParentId:      copyInfo.Authority.ParentId,
	//	Children:      []authorities.Authorities{},
	//}
	//if !authorities.RecordNotFound(g.Map{"authority_id": copyInfo.Authority.AuthorityId}) {
	//	return authority, errors.New("存在相同角色id")
	//}
	//menusReturn, err = GetMenuAuthority(copyInfo.OldAuthorityId)
	//var baseMenu []menus.Entity
	//for _, v := range menusReturn {
	//	intNum, _ := strconv.Atoi(v.MenuId)
	//	v.Entity.Id = uint(intNum)
	//	baseMenu = append(baseMenu, v.Entity)
	//}
	//copyInfo.Authority.Menus = baseMenu
	//_, err = menus.Insert(&copyInfo.Authority)
	//paths := GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	//if err = UpdateCasbin(copyInfo.Authority.AuthorityId, paths); err != nil {
	//	_ = DeleteAuthority(&copyInfo.Authority)
	//}
	//return &copyInfo.Authority, err
	return
}

// @title    UpdateAuthority
// @description   更改一个角色
func UpdateAuthority(update *request.UpdateAuthority) (authority *authorities.Entity, err error) {
	updateData := &authorities.Entity{
		AuthorityId:   update.AuthorityId,
		AuthorityName: update.AuthorityName,
		ParentId:      update.ParentId,
	}
	if _, err = authorities.Update(g.Map{"authority_id": update.AuthorityId}); err != nil {
		return updateData, err
	}
	return updateData, nil
}

// DeleteAuthority Delete role
// DeleteAuthority 删除角色
func DeleteAuthority(auth *request.DeleteAuthority) (err error) {
	if _, err = admins.FindOne(g.Map{"authority_id": auth.AuthorityId}); err != nil {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if _, err = authorities.Delete(g.Map{"parent_id": auth.AuthorityId}); err != nil {
		return errors.New("此角色存在子角色不允许删除")
	}
	if _, err = authority_menu.Delete(g.Map{"authority_id": auth.AuthorityId}); err != nil {
		return errors.New("菜单删除失败")
	}
	ClearCasbin(0, auth.AuthorityId)
	return err
}

// GetInfoList Get data by page
// GetInfoList 分页获取数据
func GetAuthorityInfoList(info *request.PageInfo) (list interface{}, total int, err error) {
	var associated []*authority_resources.Entity
	authorityList := ([]*authorities.Authorities)(nil)
	authorityDb := global.GFVA_DB.Table("authorities").Safe()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = authorityDb.Where(g.Map{"parent_id": "0"}).Limit(limit).Offset(offset).Structs(&authorityList)
	for _, v := range authorityList {
		associated, err = authority_resources.FindAll(g.Map{"authority_id": v.AuthorityId})
		for _, a := range associated {
			err = authorityDb.Where(g.Map{"authority_id": a.ResourcesId}).Scan(v.DataAuthority) // 资源权限
		}
		err = authorityDb.Where(g.Map{"parent_id": v.AuthorityId}).Scan(v.Children) // 子用户
	}
	if err != nil {
		return authorityList, total, errors.New("查询失败! ")
	}
	return authorityList, total, err
}

// GetAuthorityInfo Get all character information
// GetAuthorityInfo 获取所有角色信息
func GetAuthorityInfo(auth *authorities.Authorities) (err error, sa authorities.Authorities) {
	//err = global.GVA_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	//return err, sa
	return
}

// SetDataAuthority Set role resource permissions
// SetDataAuthority 设置角色资源权限
func SetDataAuthority(auth *authorities.Authorities) (err error) {
	condition := g.Map{
		"authority_id": auth.AuthorityId,
		"resources_id": auth.ResourcesId,
	}
	for _, v := range auth.DataAuthority {
		condition["resources_id"] = v.ResourcesId
		if _, err = authority_resources.Insert(condition); err != nil {
			return err
		}
	}
	return
}

// SetMenuAuthority Menu and character binding
// SetMenuAuthority 菜单与角色绑定
func SetMenuAuthority(auth *authorities.Authorities) (err error) {
	//var s model.SysAuthority
	//global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	//err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus).Error
	//return err
	return
}
