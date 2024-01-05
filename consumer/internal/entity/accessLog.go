package entity

import "time"

type AccessLog struct {
	Timestamp    time.Time
	IsNormalMode bool
	IsLogin      bool
	UserName     string
	DeviceName   string
	Ip           string
}
