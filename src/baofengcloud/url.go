package baofengcloud

import (
	"errors"
	"net/url"
	"strings"
	"time"
)

func BuildSwfPlayUrl(conf *Configure, fileType FileType, videoUrl, swfUrl string, timeOut time.Duration) (string, error) {
	fid := ""
	for _, p := range strings.Split(videoUrl, "&") {
		kv := strings.Split(p, "=")
		if len(kv) == 2 && kv[0] == "fid" {
			fid = kv[1]
			break
		}
	}

	if len(fid) == 0 {
		return "", errors.New("invalid vidoe url, fid not found!")
	}

	if len(swfUrl) == 0 {
		swfUrl = SwfUrl
	}

	playUrl := url.URL{}
	playUrl.Path = swfUrl

	params := url.Values{}
	params.Add("vk", videoUrl)

	if fileType == Private {
		playToken := CreatePlayToken(conf.AccessKey, conf.SecretKey, fid, timeOut)

		params.Add("tk", playToken)
	}

	playUrl.RawQuery = params.Encode()

	return playUrl.String(), nil
}
