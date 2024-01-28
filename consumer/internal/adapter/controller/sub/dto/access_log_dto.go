package dto

import "time"

type AccessLogDto struct {
	Timestamp   time.Time
	IsNormalMode bool
	IsLogin     bool
	UserName    string
	DeviceName  string
	Ip          string
}
