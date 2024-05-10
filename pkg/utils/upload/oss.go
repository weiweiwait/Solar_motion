package upload

import (
	"Solar_motion/config"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

func ToQiNiu(file multipart.File, fileSize int64) (url string, err error) {
	qConfig := config.Config.Oss
	AccessKey := qConfig.AccessKeyId
	SerectKey := qConfig.AccessKeySecret
	Bucket := qConfig.BucketName
	ImgUrl := qConfig.QiNiuServer
	mac := qbox.NewMac(AccessKey, SerectKey)
	putPolicy := storage.PutPolicy{
		Scope:   Bucket,
		Expires: 7200,
	}
	uploadToken := putPolicy.UploadToken(mac)
	// 上传Config对象
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, uploadToken, file, fileSize, &putExtra)
	if err != nil {
		println(err)
		return "", err
	}

	url = ImgUrl + ret.Key
	return url, nil
}
