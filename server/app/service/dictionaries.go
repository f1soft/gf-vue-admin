package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model/dictionaries"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"
)

// CreateDictionary create a Dictionary
// CreateDictionary 创建Dictionary
func CreateDictionary(create *request.CreateDictionary) (err error) {
	if _, err = dictionaries.FindOne(g.Map{"type": create.Type}); err != nil {
		insert := &dictionaries.Entity{
			Name:   create.Name,
			Type:   create.Name,
			Status: utils.BoolToInt(create.Status),
			Desc:   create.Name,
		}
		if _, err = dictionaries.Insert(insert); err != nil {
			return errors.New("创建Dictionary失败")
		}
		return nil
	}
	return errors.New("存在相同的type，不允许创建")
}

// DeleteDictionary delete a Dictionary
// DeleteDictionary 删除Dictionary
func DeleteDictionary(delete *request.DeleteDictionary) (err error) {
	if _, err = dictionaries.Delete(g.Map{"id": delete.Id}); err != nil {
		return errors.New("删除Dictionary失败")
	}
	// TODO 删除DictionaryDetails
	return
}

// UpdateDictionary update Dictionary
// UpdateDictionary 更新 Dictionary
func UpdateDictionary(update *request.UpdateDictionary) (err error) {
	var dictFromDb *dictionaries.Entity
	condition := g.Map{"id": update.Id}
	updateData := g.Map{
		"name":   update.Name,
		"type":   update.Type,
		"status": update.Status,
		"desc":   update.Desc,
	}
	if dictFromDb, err = dictionaries.FindOne(condition); err != nil {
		return errors.New("记录不存在,更新失败")
	}
	if dictFromDb.Type == update.Type || dictionaries.RecordNotFound(g.Map{"type": update.Type}) {
		_, err = dictionaries.Update(updateData, condition)
		return err
	}
	return errors.New("更新失败")
}

// FindDictionary Find a Dictionary with id
// FindDictionary 用id查询Dictionary
func FindDictionary(find *request.FindDictionary) (dictReturn *dictionaries.Entity, err error) {
	return dictionaries.FindOne("id = ? OR type = ?", find.Id, find.Type)
}

// GetDictionaryInfoList get Dictionary list by pagination
// GetDictionaryInfoList 通过分页获得Dictionary列表
func GetDictionaryInfoList(info *request.DictionaryInfoList) (list interface{}, total int, err error) {
	var dictionaryList []*dictionaries.Entity
	condition := g.Map{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.DB("default").Table("dictionaries").Safe()
	if info.Name != "" {
		condition["name like ?"] = "%" + info.Name + "%"
	}
	if info.Type != "" {
		condition["type like ?"] = "%" + info.Type + "%"
	}
	if info.Type != "" {
		condition["status"] = info.Status
	}
	if info.Desc != "" {
		condition["desc like ?"] = "%" + info.Desc + "%"
	}
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Scan(&dictionaryList)
	return dictionaryList, total, err
}
