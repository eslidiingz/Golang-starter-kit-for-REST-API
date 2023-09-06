package models

import (
	"gorm.io/gorm"
)

type Livestream struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserId    uint   `json:"user_id" gorm:"index"`
	Name      string `json:"name"`
	StreamKey string `json:"stream_key"`
	StreamUrl string `json:"stream_url"`
	M3u8Url   string `json:"m3u8_url"`
	Response  string `json:"response"`
	Note      string `json:"note"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
	gorm.Model
}
