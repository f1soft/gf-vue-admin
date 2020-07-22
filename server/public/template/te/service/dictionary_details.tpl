package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model/dictionary_details"
	"server/library/utils"
	"github.com/gogf/gf/frame/g"
)

// CreateDictionaryDetail create a DictionaryDetail
// CreateDictionaryDetail 创建一个DictionaryDetail
func CreateDictionaryDetail(create *request.CreateDictionaryDetail) (err error) {
	insert := &dictionary_details.Entity{
		Label:        create.Label,
		Value:        create.Value,
		Status:       utils.BoolToInt(create.Status),
		Sort:         create.DictionaryId,
		DictionaryId: create.DictionaryId,
	}
	_, err = dictionary_details.Insert(insert)
	return nil
}

// CreateDictionaryDetail create a DictionaryDetail
// CreateDictionaryDetail 创建一个DictionaryDetail
func DeleteDictionaryDetail(delete *request.DeleteDictionaryDetail) (err error) {
	_, err = dictionary_details.Delete(g.Map{"id": delete.Id})
	return err
}

// UpdateDictionaryDetail Update a DictionaryDetail
// UpdateDictionaryDetail 更新一个DictionaryDetail
func UpdateDictionaryDetail(update *request.UpdateDictionaryDetail) (err error) {
	condition := g.Map{"id": update.Id}
	updateData := g.Map{
		"label":         update.Label,
		"value":         update.Value,
		"status":        update.Status,
		"sort":          update.Sort,
		"dictionary_id": update.DictionaryId,
	}
	if _, err = dictionary_details.FindOne(condition); err != nil {
		return errors.New("记录不存在, 更新失败")
	}
	if _, err = dictionary_details.Update(updateData, condition); err != nil {
		return errors.New("更新失败")
	}
	return
}

// FindDictionaryDetail Query DictionaryDetail with id
// FindDictionaryDetail 用id查询DictionaryDetail
func FindDictionaryDetail(find *request.FindDictionaryDetail) (dataReturn *dictionary_details.Entity, err error) {
	return dictionary_details.FindOne(g.Map{"id": find.Id})
}

// GetDictionaryDetailList Paging to get a list of DictionaryDetails
// GetDictionaryDetailList 分页获取DictionaryDetail列表
func GetDictionaryDetailList(get *request.GetDictionaryDetailList) (list interface{}, total int, err error) {
	var dictionaryList []*dictionary_details.Entity
	condition := g.Map{}
	limit := get.PageSize
	offset := get.PageSize * (get.Page - 1)
	db := g.DB("default").Table("dictionaries").Safe()
	if get.Label != "" {
		condition["label like ?"] = "%" + get.Label + "%"
	}
	if get.Value != 0 {
		condition["value"] = get.Value
	}
	if get.Status {
		condition["status"] = get.Status
	}
	if get.DictionaryId != 0 {
		condition["dictionary_id"] = get.DictionaryId
	}
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Scan(&dictionaryList)
	return dictionaryList, total, err
}