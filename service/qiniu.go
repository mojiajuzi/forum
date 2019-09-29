package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mojiajuzi/forum/config"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

//QiniuSend 七牛发送对象
type QiniuSend struct {
	Ak       string
	Sk       string
	Mac      *qbox.Mac
	Token    string
	Bucket   string
	Cfg      storage.Config
	Uploader *storage.FormUploader
	Ret      MyPutRet
	PutExtra storage.PutExtra
	BaseURL  string
}

//MyPutRet 解析结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

//Init 初始化七牛云上传对象
func (q *QiniuSend) Init() {
	putPolicy := storage.PutPolicy{
		Scope:      config.Config("QINIU_BUCKET", "image"),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires = 7200
	accessKey := config.Config("QINIU_ACCESSKEY", "access_key")
	secretKey := config.Config("QINIU_SECRETKEY", "secret_key")
	q.Mac = qbox.NewMac(accessKey, secretKey)
	q.Token = putPolicy.UploadToken(q.Mac)

	q.BaseURL = config.Config("QINIU_BASE_URL", "base_url")

	//初始化配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = true
	cfg.UseCdnDomains = true
	q.Cfg = cfg

	//上传表单初始化
	q.Uploader = storage.NewFormUploader(&q.Cfg)
	q.Ret = MyPutRet{}
	q.PutExtra = storage.PutExtra{}
}

//UploadByURL 通过地址获取文件信息并上传文件到七牛云，并返回对应的文件名
func (q *QiniuSend) UploadByURL(fileURL string) (string, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("请求返回状态错误,响应状态码为:code, %d", resp.StatusCode)
		return "", errors.New(msg)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	dataLen := int64(len(data))
	key := q.genereteFileName(fileURL)

	q.PutExtra.Params = map[string]string{
		"x:name": key,
	}
	err = q.Uploader.Put(context.Background(), &q.Ret, q.Token, key, bytes.NewReader(data), dataLen, &q.PutExtra)
	if err != nil {
		return "", err
	}
	finalURL := q.BaseURL + q.PutExtra.Params["x:name"]
	return finalURL, nil
}

//genereteFileName 生成新的文件名
func (q *QiniuSend) genereteFileName(name string) string {
	now := time.Now().String()
	name = now + name
	h := md5.New()
	h.Write([]byte(name))
	cipherStr := h.Sum(nil)
	prefix := "/image/"
	return prefix + hex.EncodeToString(cipherStr)
}

//UploadByBody 上传文件结构体
func (q *QiniuSend) UploadByBody(fileName string, size int64, r *bytes.Reader) (string, error) {
	key := q.genereteFileName(fileName)

	q.PutExtra.Params = map[string]string{
		"x:name": key,
	}
	err := q.Uploader.Put(context.Background(), &q.Ret, q.Token, key, r, size, &q.PutExtra)
	if err != nil {
		return "", err
	}
	finalURL := q.BaseURL + q.PutExtra.Params["x:name"]
	return finalURL, nil
}
