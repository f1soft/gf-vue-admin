package service

import (
	"errors"
	"gf-server/app/api/request"
	"gf-server/app/api/response"
	"gf-server/app/model"
	"gf-server/app/model/authorities"
	"gf-server/app/model/menus"
	"gf-server/library/global"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"strconv"
)

// CreateAuthority Create a role
// CreateAuthority 创建一个角色
func CreateAuthority(auth request.CreateAuthority) (authority *authorities.Entity, err error) {
	authority = &authorities.Entity{
		AuthorityId:   auth.AuthorityId,
		AuthorityName: auth.AuthorityName,
		ParentId:      auth.ParentId,
	}
	if !authorities.RecordNotFound(g.Map{"authority_id": auth.AuthorityId}) {
		return authority, errors.New("存在相同角色id")
	}
	if _, err = authorities.Insert(authority); err != nil {
		return authority, errors.New("创建角色失败")
	}
	return authority, nil
}

// CopyAuthority Copy a character
// CopyAuthority 复制一个角色
func CopyAuthority(copyInfo response.AuthorityCopyResponse) (authority *authorities.Authorities, err error) {
	var menusReturn []model.AuthorityMenus
	authority = &authorities.Authorities{
		AuthorityId:   copyInfo.Authority.AuthorityId,
		AuthorityName: copyInfo.Authority.AuthorityName,
		ParentId:      copyInfo.Authority.ParentId,
		Children:      []authorities.Authorities{},
	}
	if !authorities.RecordNotFound(g.Map{"authority_id": copyInfo.Authority.AuthorityId}) {
		return authority, errors.New("存在相同角色id")
	}
	menusReturn, err = GetMenuAuthority(copyInfo.OldAuthorityId)
	var baseMenu []menus.Entity
	for _, v := range menusReturn {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.Entity.Id = uint(intNum)
		baseMenu = append(baseMenu, v.Entity)
	}
	copyInfo.Authority.Menus = baseMenu
	_, err = menus.Insert(&copyInfo.Authority)
	paths := GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	if err = UpdateCasbin(copyInfo.Authority.AuthorityId, paths); err != nil {
		_ = DeleteAuthority(&copyInfo.Authority)
	}
	return &copyInfo.Authority, err
}

// @title    UpdateAuthority
// @description   更改一个角色
func UpdateAuthority(auth *request.UpdateAuthority) (authority *request.UpdateAuthority, err error) {
	_, err = authorities.Update(auth, "authority_id", auth.AuthorityId)
	return auth, err
}

// DeleteAuthority Delete role
// DeleteAuthority 删除角色
func DeleteAuthority(auth *authorities.Authorities) (err error) {
	if _, err = authorities.Delete(g.Map{"authority_id": auth.AuthorityId}); err != nil {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if _, err = authorities.Delete(g.Map{"parent_id": auth.AuthorityId}); err != nil {
		return errors.New("此角色存在子角色不允许删除")
	}
	// TODO 关联
	//db := global.GVA_DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth)
	//if len(auth.SysBaseMenus) > 0 {
	//	err = db.Association("SysBaseMenus").Delete(auth.SysBaseMenus).Error
	//} else {
	//	err = db.Error
	//}
	ClearCasbin(0, auth.AuthorityId)
	return err
}

// GetInfoList Get data by page
// GetInfoList 分页获取数据
func GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int, err error) {
	var (
		authorityList []authorities.Authorities
		dataAuthority []*authorities.Authorities
		dataEntity    []*authorities.Entity
	)
	db := global.GFVA_DB.Table("authorities").Safe()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	if err = db.Where(g.Map{"parent_id": "0"}).Limit(limit).Offset(offset).Scan(&authorityList);err != nil {
		return authorityList, total, errors.New("查询失败! ")
	}
	// TODO 自关联
	if dataEntity, err = authorities.FindAll(); err != nil{
		return authorityList, total, errors.New("查询失败! ")
	}
	if len(authorityList) > 0 {
		for k := range authorityList {
			_ = gconv.Structs(dataEntity, dataAuthority)
			authorityList[k].Children = dataAuthority
			err = findChildrenAuthority(&authorityList[k])
		}
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
	//var s model.SysAuthority
	//global.GVA_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	//err := global.GVA_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId).Error
	//return
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

// findChildrenAuthority Query subrole
// findChildrenAuthority 查询子角色
func findChildrenAuthority(authority *authorities.Authorities) (err error) {
	//err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	//if len(authority.Children) > 0 {
	//	for k := range authority.Children {
	//		err = findChildrenAuthority(&authority.Children[k])
	//	}
	//}
	//return err
	return
}
