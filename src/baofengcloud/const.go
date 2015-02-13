package baofengcloud

import (
	"time"
)

const (
	Paas = 0
	Saas = 1
)

type ServiceType int

const (
	Private = 0
	Public  = 1
)

type FileType int

const (
	Full    = 0
	Partial = 1
)

type UploadType int

const (
	TokenTimeout     = 1 * time.Hour
	ploadBlockSize   = 4 * 1024 * 1024
	UploadRequestUrl = "http://access.baofengcloud.com/upload"
	DeleteRequestUrl = "http://access.baofengcloud.com/delete"
	QueryRequestUrl  = "http://access.baofengcloud.com/query"
	UpdateRequestUrl = "http://access.baofengcloud.com/change"
	SwfUrl           = "http://www.baofengcloud.com/html/swf/player/cloud.swf"
)
