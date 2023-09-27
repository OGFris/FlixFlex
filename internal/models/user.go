package models

type User struct {
	Model
	Email    string `gorm:"not null;unique"`
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Disabled bool   `gorm:"not null;default:false"`
}
