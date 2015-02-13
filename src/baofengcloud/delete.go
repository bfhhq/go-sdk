package baofengcloud

import (
	"encoding/json"
)

func DeleteFile(conf *Configure, serviceType ServiceType, fileName, fileKey, callbackUrl string) (*Result, error) {

	token := CreateDeleteToken(conf.AccessKey, conf.SecretKey, serviceType, fileName, fileKey, TokenTimeout, callbackUrl)

	return DeleteFileByToken(token)
}

func DeleteFileByToken(token string) (*Result, error) {

	body, err := PostToken(DeleteRequestUrl, token)
	if err != nil {
		return nil, err
	}

	result := &Result{}
	err = json.Unmarshal(body, result)

	return result, err
}
