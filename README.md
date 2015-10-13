# Baofeng Cloud SDK for Golang

安装
``` bash
go get github.com/baofengcloud/go-sdk/src/baofengcloud
```
使用
``` go
import "github.com/baofengcloud/go-sdk/src/baofengcloud"
```
配置AK/SK
``` go
var conf = Configure{
	AccessKey: "",
	SecretKey: "",
}
```
上传
``` go
baofengcloud.UploadFile2(&conf, baofengcloud.Paas, baofengcloud.Public, "C:\\test.mp4", "test.mp4", "", "")
```
查询
``` go
baofengcloud.QueryFile(&conf, baofengcloud.Paas, "test.mp4", "")
```
删除
``` go
baofengcloud.DeleteFile(&conf, baofengcloud.Paas, "test.mp4", "", "")
```

# 关于

基于 [暴风云视频API](http://www.baofengcloud.com/apisdk/doc.html) 构建。
