package testcase

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestDownIcon(t *testing.T) {
	imagePath := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
	savePath := "F:\\go\\src\\test\\static\\"
	resp, _ := http.Get(imagePath)
	body, _ := ioutil.ReadAll(resp.Body)
	name := UniqueId()
	out, _ := os.Create(savePath + name + ".png")
	io.Copy(out, bytes.NewReader(body))
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func TestAliOssUpload(t *testing.T) {
	imagePath := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
	saveUrl := AliOssUpload(imagePath)
	if saveUrl==imagePath{
		t.Error("ali oss upload error",imagePath)
	}
}

func AliOssUpload(url string) string {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。

	//doc https://help.aliyun.com/document_detail/88601.html

	client, err := oss.New("yourEndpoint", "yourAccessKeyId", "yourAccessKeySecret")
	if err != nil {
		//阿里云初始化失败
		return url
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("examplebucket")

	if err != nil {
		//bucket error
		return url
	}
	env := "local"
	filename := UniqueId()
	savePath := "app/" + env + "/icon/" + filename + ".png"
	// 指定待上传的网络流。
	res, _ := http.Get(url)

	if err != nil {
		return url
	}
	err = bucket.PutObject(savePath, io.Reader(res.Body))
	if err != nil {
		//阿里云上传文件失败
	} else {
		url = savePath
	}
	return url
}
