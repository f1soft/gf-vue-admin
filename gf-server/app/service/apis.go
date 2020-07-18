package service

import (
	"errors"
	"gf-server/app/api/request"
	"gf-server/app/model/apis"
	"github.com/gogf/gf/util/gconv"

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
	if err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method); err != nil{
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
func GetAllApis() (apisReturn []*apis.Entity, err error) {
	return apis.FindAll()
}

// GetApiInfoList 分页获取数据
// 未测试,不确定逻辑是否走得通
func GetApiInfoList(api *request.GetApiList) (list interface{}, total int, err error) {
	db := g.DB("default")
	limit := api.PageSize
	offset := api.PageSize * (api.Page - 1)
	condition := g.Map{}
	var apisList []apis.Entity
	if api.Path != "" {
		condition["path like ?"] = "%" + api.Path + "%"
	}
	if api.Description != "" {
		condition["description like ?"] = "%" + api.Description + "%"
	}
	if api.Method != "" {
		condition["method"] = api.Method
	}
	if api.ApiGroup != "" {
		condition["api_group"] = api.ApiGroup
	}
	if total, findErr := db.Table("apis").Where(condition).Count(); findErr != nil {
		return apisList, total, findErr
	}
	if api.OrderKey != "" && api.Desc == true {
		orderStr := api.OrderKey + " desc"
		r, err := db.Table("apis").Where(condition).Order(orderStr).Limit(limit).Offset(offset).FindAll()
		if err != nil {
			return apisList, total, err
		}
		if err := gconv.Struct(r, &apisList); err != nil {
			return apisList, total, err
		}
	}
	if api.OrderKey != "" && api.Desc == false {
		r, err := db.Table("apis").Where(condition).Order("api_group").Limit(limit).Offset(offset).FindAll()
		if err != nil {
			return apisList, total, err
		}
		if err := gconv.Struct(r, &apisList); err != nil {
			return apisList, total, err
		}
	}
	return apisList, total, err
}
