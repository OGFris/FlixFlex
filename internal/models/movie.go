package models

type Movie struct {
	Model
	Title     string `gorm:"not null;unique"`
	Synopsis  string `gorm:"not null"`
	PosterUri string `gorm:"not null"`
	TmdbID    uint   `gorm:"not null;unique"`
}
