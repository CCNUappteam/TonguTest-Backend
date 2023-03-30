package service

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	qiniu_storage "github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"time"
)

func UploadFile(file *multipart.FileHeader) (error, string) {
	src, err := file.Open()
	if err != nil {
		return err, ""
	}
	initOSS()

	defer src.Close()

	putPolicy := qiniu_storage.PutPolicy{
		Scope: bucketName,
	}

	mac := qbox.NewMac(accessKey, secretKey)

	// 获取上传凭证
	upToken := putPolicy.UploadToken(mac)

	// 配置参数
	cfg := qiniu_storage.Config{
		Zone:          &qiniu_storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	formUploader := qiniu_storage.NewFormUploader(&cfg)

	ret := qiniu_storage.PutRet{}        // 上传返回后的结果
	putExtra := qiniu_storage.PutExtra{} // 额外参数

	// 自定义文件名及后缀
	key := "(" + time.Now().String() + ")" + file.Filename

	if err := formUploader.Put(context.Background(), &ret,
		upToken, key, src, file.Size, &putExtra); err != nil {
		fmt.Println(err)
		return err, ""
	}

	return nil, domainName + "/" + ret.Key
}
