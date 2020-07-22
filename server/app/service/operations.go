package service

import (
	"server/app/api/request"
	"server/app/model/operations"

	"github.com/gogf/gf/frame/g"
)

// CreateOperation create a operation
// CreateOperation 创建operation
func CreateOperation(create *request.CreateOperation) (err error) {
	insert := operations.Entity{
		Ip:           create.Ip,
		Method:       create.Method,
		Path:         create.Path,
		Status:       create.Status,
		Latency:      create.Latency,
		Agent:        create.Agent,
		ErrorMessage: create.ErrorMessage,
		Request:      create.Request,
		UserId:       create.UserId,
		Response:     create.Response,
	}
	_, err = operations.Insert(&insert)
	return err
}

// DeleteOperation delete Operation
// DeleteOperation 删除 Operation
func DeleteOperation(delete *request.DeleteOperation) (err error) {
	_, err = operations.Delete(g.Map{"id": delete.Id})
	return err
}

// DeleteOperations delete Operations
// DeleteOperations 删除 Operations
func DeleteOperations(deletes *request.DeleteOperations) (err error) {
	_, err = operations.Delete(g.Map{"id IN (?)": deletes.Ids})
	return err
}

// UpdateOperation update Operations
// UpdateOperation 更新 Operations
func UpdateOperation(update *request.UpdateOperation) (err error) {
	updateData := g.Map{
		"id":            update.Id,
		"ip":            update.Ip,
		"method":        update.Method,
		"path":          update.Path,
		"status":        update.Status,
		"latency":       update.Latency,
		"agent":         update.Agent,
		"error_message": update.ErrorMessage,
		"request":       update.Request,
		"userId":        update.UserId,
		"response":      update.Response,
	}
	_, err = operations.Save(updateData)
	return err
}

// FindOperation Gets a single Operation based on id
// FindOperation 根据id获取单条Operation
func FindOperation(find *request.FindOperation) (operation *operations.Entity, err error) {
	return operations.FindOne(g.Map{"id": find.Id})
}

// GetOperationList Page out the Operation list
// GetOperationList 分页获取Operation列表
func GetOperationList(info *request.GetOperationList) (list interface{}, total int, err error) {
	var operationList []*operations.Entity
	condition := g.Map{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.DB("default").Table("operations").Safe()
	if info.Method != "" {
		condition["method"] = info.Method
	}
	if info.Path != "" {
		condition["path like ?"] = "%" + info.Method + "%"
	}
	if info.Status != 0 {
		condition["status"] = info.Status
	}
	total, err = db.Where(condition).Count()
	err = db.Order("id desc").Limit(limit).Offset(offset).Scan(&operationList)
	return operationList, total, err
}
