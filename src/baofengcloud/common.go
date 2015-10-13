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

type Representation struct {
	FileSize   int64
	Duration   int64
	Definition int
	Width      int
	Height     int
	Resolution string
	Urls       []string `json:"urllist"`
}

type FileInfo2 struct {
	UserId          int
	FileId          string
	FileName        string
	FileKey         string
	FileSize        int64
	Duration        int64
	ServiceType     ServiceType
	FileType        FileType         `json:"ifpublic"`
	Urls            []string         `json:"urllist"`
	Representations []Representation `json:"gcids"`
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

func Get(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
