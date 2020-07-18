package service

import (
	"errors"
	"gf-server/app/api/request"
	"gf-server/app/model/dictionaries"

	"github.com/gogf/gf/frame/g"
)

// CreateDictionary create a Dictionary
// CreateDictionary 创建Dictionary
func CreateDictionary(dict *request.CreateDictionary) (err error) {
	if !dictionaries.RecordNotFound(g.Map{"type": dict.Type}) {
		return errors.New("存在相同的type，不允许创建")
	}
	if _, err = dictionaries.Insert(&dict); err != nil {
		return errors.New("创建Dictionary失败")
	}
	return
}

// DeleteDictionary delete a Dictionary
// DeleteDictionary 删除Dictionary
func DeleteDictionary(dict *request.DeleteDictionary) (err error) {
	if _, err = dictionaries.Delete(g.Map{"id": dict.Id}); err != nil {
		return errors.New("删除Dictionary失败")
	}
	// TODO 删除DictionaryDetails
	return
}

// UpdateDictionary update Dictionary
// UpdateDictionary 更新 Dictionary
func UpdateDictionary(dict *request.UpdateDictionary) (err error) {
	var dictFromDb *dictionaries.Entity
	condition := g.Map{
		"name":   dict.Name,
		"type":   dict.Type,
		"status": dict.Status,
		"desc":   dict.Desc,
	}
	if dictFromDb, err = dictionaries.FindOne(g.Map{"id": dict.Id}); err != nil {
		return errors.New("更新失败")
	}
	if dictFromDb.Type == dict.Type || dictionaries.RecordNotFound(g.Map{"type": dict.Type}) {
		_, err = dictionaries.Update(condition)
		return err
	}
	return errors.New("更新失败")
}

// GetDictionary get the info of a Dictionary
// GetDictionary 获取Dictionary的信息
func GetDictionary(dict *request.FindDictionary) (dictReturn *dictionaries.Entity, err error) {
	return dictionaries.FindOne("id = ? OR type = ?", dict.Id, dict.Type)
}

// GetDictionaryInfoList get Dictionary list by pagination
// GetDictionaryInfoList 通过分页获得Dictionary列表
func GetDictionaryInfoList(info *request.DictionaryInfoList) (list interface{}, total int, err error) {
	var dictionarys []*dictionaries.Entity
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
	err = db.Limit(limit).Offset(offset).Where(condition).Scan(&dictionarys)
	return dictionarys, total, err
}
