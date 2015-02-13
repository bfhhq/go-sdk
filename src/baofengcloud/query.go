package baofengcloud

import (
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
