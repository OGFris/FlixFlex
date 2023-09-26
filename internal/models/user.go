package models

type User struct {
	Model
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Disabled bool   `gorm:"not null;default:false"`
}
