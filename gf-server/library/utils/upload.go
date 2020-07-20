package utils

import (
	"context"
	"fmt"
	"gf-server/library/global"
	"mime/multipart"
	"time"

	"github.com/gogf/gf/frame/g"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// 接收两个参数 一个文件流 一个 bucket 你的七牛云标准空间的名字
func Upload(file *multipart.FileHeader) (path string, key string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: g.Cfg().GetString("qiniu.Bucket"),
	}
	mac := qbox.NewMac(g.Cfg().GetString("qiniu.AccessKey"), g.Cfg().GetString("qiniu.SecretKey"))
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	f, e := file.Open()
	if e != nil {
		fmt.Println(e)
		return "", "", e
	}
	dataLen := file.Size
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	err = formUploader.Put(context.Background(), &ret, upToken, fileKey, f, dataLen, &putExtra)
	if err != nil {
		global.GFVA_LOG.Error("upload file fail:", err)
		return "", "", err
	}
	return g.Cfg().GetString("qiniu.ImgPath") + "/" + ret.Key, ret.Key, err
}

func DeleteFile(key string) error {
	mac := qbox.NewMac(g.Cfg().GetString("qiniu.AccessKey"), g.Cfg().GetString("qiniu.SecretKey"))
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	// cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(g.Cfg().GetString("qiniu.AccessKey"), key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
