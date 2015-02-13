package baofengcloud

import (
	"encoding/json"
)

func UpdateFile(conf *Configure, serviceType ServiceType, fileName, fileKey string, fileType FileType) (*FileInfo, error) {

	token := CreateUpdateToken(conf.AccessKey, conf.SecretKey, serviceType, fileType, fileName, fileKey)

	return UpdateFileByToken(token)
}

func UpdateFileByToken(token string) (*FileInfo, error) {

	body, err := PostToken(UpdateRequestUrl, token)
	if err != nil {
		return nil, err
	}

	fileInfo := &FileInfo{}
	err = json.Unmarshal(body, fileInfo)

	return fileInfo, err
}
