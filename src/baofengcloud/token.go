package baofengcloud

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

func createToken(msg []byte, accessKey, secretKey string) string {

	encodedMsg := base64.StdEncoding.EncodeToString(msg)

	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(encodedMsg))
	sign := mac.Sum(nil)
	encodedSign := base64.StdEncoding.EncodeToString(sign)

	token := accessKey + ":" + encodedSign + ":" + encodedMsg

	return token
}

func deadline(timeOut time.Duration) int64 {
	return time.Now().UTC().Unix() + int64(timeOut.Seconds())
}

func createJsonToken(data map[string]interface{}, accessKey, secretKey string, timeOut time.Duration) string {

	if timeOut > 0 {
		data["deadline"] = deadline(timeOut)
	}

	jsonStr, _ := json.Marshal(data)

	return createToken(jsonStr, accessKey, secretKey)
}

func createUrlToken(url, accessKey, secretKey string, timeOut time.Duration) string {

	newUrl := fmt.Sprintf("%s&deadline=%d", url, deadline(timeOut))

	return createToken([]byte(newUrl), accessKey, secretKey)
}

func CreateUploadToken(accessKey, secretKey string, serviceType ServiceType,
	fileType FileType, uploadType UploadType, fileName, fileKey string,
	fileSize int64, timeOut time.Duration, callbackUrl string) string {

	data := make(map[string]interface{})

	data["uptype"] = uploadType
	data["servicetype"] = serviceType
	data["filekey"] = fileKey
	data["filename"] = fileName
	data["filesize"] = fileSize
	data["filetype"] = fileType
	data["callbackurl"] = callbackUrl

	return createJsonToken(data, accessKey, secretKey, timeOut)
}

func CreateDeleteToken(accessKey, secretKey string, serviceType ServiceType,
	fileName, fileKey string, timeOut time.Duration, callbackUrl string) string {

	data := make(map[string]interface{})

	data["servicetype"] = serviceType
	data["filekey"] = fileKey
	data["filename"] = fileName
	data["callbackurl"] = callbackUrl

	return createJsonToken(data, accessKey, secretKey, timeOut)
}

func CreateQueryToken(accessKey, secretKey string, serviceType ServiceType, fileName, fileKey string) string {

	data := make(map[string]interface{})

	data["servicetype"] = serviceType
	data["filekey"] = fileKey
	data["filename"] = fileName

	return createJsonToken(data, accessKey, secretKey, 0)
}

func CreateUpdateToken(accessKey, secretKey string, serviceType ServiceType, fileType FileType, fileName, fileKey string) string {

	data := make(map[string]interface{})

	data["servicetype"] = serviceType
	data["filekey"] = fileKey
	data["filename"] = fileName
	data["filetype"] = fileType

	return createJsonToken(data, accessKey, secretKey, 0)
}

func CreatePlayToken(accessKey, secretKey, fid string, timeOut time.Duration) string {

	url := "id=" + fid

	return createUrlToken(url, accessKey, secretKey, timeOut)
}

func CreateDownloadToken(accessKey, secretKey, fid string, timeOut time.Duration) string {

	url := "id=" + fid

	return createUrlToken(url, accessKey, secretKey, timeOut)
}
