package dto

import "time"

type AccessLogDto struct {
	Timestamp    time.Time `json:"timestamp" binding:"required"`
	IsNormalMode *bool     `json:"is_normal_mode" binding:"required"`
	IsLogin      *bool     `json:"is_login" binding:"required"`
	UserName     string    `json:"user_name" binding:"required"`
	DeviceName   string    `json:"device_name" binding:"required"`
	Ip           string    `json:"ip" binding:"required"`
}
