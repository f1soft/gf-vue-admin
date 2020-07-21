package service

import (
	"errors"
	"gf-server/app/api/request"
	"gf-server/app/model/files"

	"github.com/gogf/gf/frame/g"
)

// UploadFile Create file upload records
// UploadFile 创建文件上传记录
func UploadFile(file *files.Entity) (err error) {
	_, err = files.Insert(file)
	return err
}

// DeleteFile Delete file records
// DeleteFile 删除文件记录
func DeleteFile(d *request.DeleteFile) error {
	if _, err := files.FindOne(g.Map{"id": d.Id}); err != nil {
		return errors.New("文件记录不存在,删除失败")
	}
	_, err := files.Delete(g.Map{"id": d.Id})
	return err
}

// GetFileList Paging fetch data
// GetFileList 分页获取数据
func GetFileList(info *request.PageInfo) (list interface{}, total int, err error) {
	var fileList []*files.Entity
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.DB("default").Table("files").Safe()
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Scan(&fileList)
	return fileList, total, err
}
