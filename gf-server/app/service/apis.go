package service

import (
	"errors"
	"gf-server/app/api/request"
	"gf-server/app/model/apis"

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

// DeleteApi 删除api
func DeleteApi(api *request.DeleteApi) error {
	_, err := apis.Delete(g.Map{"id": api.ID})
	// 	TODO ClearCasbin(1, api.Path, api.Method)
	return err
}

// UpdateApi 更新api信息
func UpdateApi(api *request.UpdateApi) error {
	oldApi, err := apis.FindOne(g.Map{"id": api.ID})
	if oldApi.Path != api.Path || oldApi.Method != api.Method {
		if !apis.RecordNotFound(g.Map{"path": api.Path, "method": api.Method}) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	}
	// TODO 		if err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method); err != nil{ return err }
	_, err = apis.Save(api)
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
