package baofengcloud

import (
	"encoding/base64"
	"encoding/json"
)

func QueryFile(conf *Configure, serviceType ServiceType, fileName, fileKey string) (*FileInfo, error) {

	token := CreateQueryToken(conf.AccessKey, conf.SecretKey, serviceType, fileName, fileKey)

	return QueryFileByToken(token)
}

func QueryFileByToken(token string) (*FileInfo, error) {

	body, err := PostToken(QueryRequestUrl, token)
	if err != nil {
		return nil, err
	}

	fileInfo := &FileInfo{}
	err = json.Unmarshal(body, fileInfo)

	return fileInfo, err
}

func QueryCdn(conf *Configure, url, fid string) (*FileInfo2, error) {
	token := CreateDownloadToken(conf.AccessKey, conf.SecretKey, fid, 0)

	query := CdnQueryUrl + "/" + base64.StdEncoding.EncodeToString([]byte(url)) + "?tk=" + token

	body, err := Get(query)
	if err != nil {
		return nil, err
	}

	fileInfo := &FileInfo2{}
	err = json.Unmarshal(body, fileInfo)

	return fileInfo, err
}
