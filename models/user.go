package models

import "gorm.io/gorm"

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Username    string `json:"username" gorm:"uniqueIndex"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	ProjectName string `json:"project_name"`
	ProjectSlug string `json:"project_slug"`
	SecretKey   string `json:"secret_key"`
	ApiKey      string `json:"api_key"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
	IsVerify    bool   `json:"is_verify" gorm:"default:true"`
	Note        string `json:"note"`
	gorm.Model
}
