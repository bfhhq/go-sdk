package baofengcloud

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func UploadFile(conf *Configure, serviceType ServiceType, fileType FileType,
	localFilePath, fileName, fileKey, callbackUrl string) error {

	f, err := os.Open(localFilePath)
	if err != nil {
		return err
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	fileSize := fi.Size()

	token := CreateUploadToken(conf.AccessKey, conf.SecretKey, serviceType,
		fileType, Full, fileName, fileKey, fileSize,
		TokenTimeout, callbackUrl)

	body, err := PostToken(UploadRequestUrl, token)
	if err != nil {
		return err
	}

	result := &Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	if result.Status != 0 {
		return fmt.Errorf("status:%d errmsg:%s", result.Status, result.ErrMsg)
	}

	uploadUrl := result.Url

	hash := md5.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		return err
	}

	hashValue := md5.Sum(nil)
	hashHex := hex.EncodeToString(hashValue[:md5.Size])

	f.Seek(0, 0)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", uploadUrl, f)
	req.Header.Add("Content-MD5", hashHex)

	//disable chunked encoding
	req.ContentLength = fileSize

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func UploadFile2(conf *Configure, serviceType ServiceType, fileType FileType,
	localFilePath, fileName, fileKey, callbackUrl string) error {

	f, err := os.Open(localFilePath)
	if err != nil {
		return err
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	fileSize := fi.Size()

	token := CreateUploadToken(conf.AccessKey, conf.SecretKey, serviceType,
		fileType, Partial, fileName, fileKey, fileSize,
		TokenTimeout, callbackUrl)

	body, err := PostToken(UploadRequestUrl, token)
	if err != nil {
		return err
	}

	result := &Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	if result.Status != 0 {
		return fmt.Errorf("status:%d errmsg:%s", result.Status, result.ErrMsg)
	}

	uploadUrl := result.Url
	//fmt.Println(uploadUrl)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", uploadUrl, nil)
	req.Header.Add("Content-Range", fmt.Sprintf("bytes */%d", fileSize))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 308 {
		return fmt.Errorf("http status code is not %d != 308", resp.StatusCode)
	}

	//TODO: to support multi-ranges as [0-100, 200-400, ....]
	startPos := int64(0)
	xranges := resp.Header.Get("Range")
	if i := strings.LastIndex(xranges, "-"); i >= 0 {
		startPos, err = strconv.ParseInt(xranges[i+1:], 0, 64)
	}

	blockLength := int64(ploadBlockSize)
	buf := make([]byte, blockLength)

	for ; startPos < fileSize; startPos += blockLength {

		if startPos+blockLength > fileSize {
			blockLength = fileSize - startPos
		}

		f.Seek(startPos, 0)
		readed, err := f.Read(buf)
		if err != nil {
			return err
		} else if blockLength != int64(readed) {
			return fmt.Errorf("io error!")
		}

		hash := md5.New()
		hash.Write(buf[0:blockLength])
		hashValue := hash.Sum(nil)
		hashHex := hex.EncodeToString(hashValue[:md5.Size])

		req, _ := http.NewRequest("POST", uploadUrl, bytes.NewReader(buf[0:blockLength]))
		req.Header.Add("Content-Range", fmt.Sprintf("bytes %d-%d/%d", startPos, startPos+blockLength-1, fileSize))
		req.Header.Add("Content-MD5", hashHex)

		//disable chunked encoding
		req.ContentLength = blockLength

		resp, err = client.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode == 200 {
			return nil
		} else if resp.StatusCode != 206 {
			return fmt.Errorf("http status code error! %d ", resp.StatusCode)
		}
	}

	return nil
}
