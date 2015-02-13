package baofengcloud

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Result struct {
	Status int
	ErrMsg string
	Url    string
}

type FileInfo struct {
	Status      int
	FileId      string
	FileName    string
	FileKey     string
	FileSize    int64
	Duration    int64
	ServiceType ServiceType
	FileType    FileType `json:"ifpublic"`
	Url         string
}

func PostToken(url, token string) ([]byte, error) {
	content := "{\"token\":\"" + token + "\"}"

	resp, err := http.Post(url, "application/json", strings.NewReader(content))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
