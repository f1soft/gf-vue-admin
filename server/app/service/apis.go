package service

import (
	"errors"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/model/apis"

	"github.com/gogf/gf/frame/g"
)

// CreateApi 创建api
func CreateApi(api *request.CreateApi) error {
	_, err := apis.FindOne(g.Map{"path": api.Path, "method": api.Method})
	if err != nil {
		return errors.New("存在相同api")
	}
	apiToDb := &apis.Entity{
		Path:        api.Path,
		Description: api.Description,
		ApiGroup:    api.ApiGroup,
		Method:      api.Method,
	}
	if _, err := apis.Insert(apiToDb); err != nil {
		return errors.New("创建api失败")
	}
	return nil
}

// UpdateApi 更新api信息
func UpdateApi(api *request.UpdateApi) error {
	oldA, err := apis.FindOne(g.Map{"id": api.ID})
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !apis.RecordNotFound(g.Map{"path": api.Path, "method": api.Method}) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	}
	if err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method); err != nil {
		return err
	}
	_, err = apis.Save(api)
	return err
}

// DeleteApi 删除api
func DeleteApi(api *request.DeleteApi) error {
	_, err := apis.Delete(g.Map{"id": api.ID})
	ClearCasbin(1, api.Path, api.Method)
	return err
}

// GetApiById 根据id获取api
func GetApiById(api *request.GetApiById) (apisReturn *apis.Entity, err error) {
	return apis.FindOne(g.Map{"id": api.ID})
}

// GetAllApis 获取所有的Api
func GetAllApis() (list []*response.Apis, err error) {
	list = ([]*response.Apis)(nil)
	db := g.DB("default").Table("apis").Safe()
	err = db.Structs(&list)
	return list, err
}

// GetApiInfoList Page to get the data
// GetApiInfoList 分页获取数据
func GetApiInfoList(api *request.GetApiList) (list []*response.Apis, total int, err error) {
	list = ([]*response.Apis)(nil)
	db := g.DB("default").Table("apis").Safe()
	limit := api.PageSize
	offset := api.PageSize * (api.Page - 1)
	condition := g.Map{}
	switch {
	case api.Path != "":
		condition["path like ?"] = "%" + api.Path + "%"
	case api.Description != "":
		condition["description like ?"] = "%" + api.Description + "%"
	case api.Method != "":
		condition["method"] = api.Method
	case api.ApiGroup != "":
		condition["api_group"] = api.ApiGroup
	}
	total, err = db.Where(condition).Count()
	if api.OrderKey != "" && api.Desc == true {
		orderStr := api.OrderKey + " desc"
		err = db.Limit(limit).Offset(offset).Order(orderStr).Where(condition).Structs(&list)
		return list, total, err
	}
	if api.OrderKey != "" && api.Desc == false {
		err = db.Where(condition).Order("api_group").Limit(limit).Offset(offset).Structs(&list)
		return list, total, err
	}
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&list)
	return list, total, err
}
