package models

import (
	"github.com/jinzhu/gorm"
)

// Gallery is our image container resources that visitors view
type Gallery struct {
	// type Model struct {
	//	ID        uint `gorm:"primary_key"`
	//	CreatedAt time.Time
	//	UpdatedAt time.Time
	//	DeletedAt *time.Time `sql:"index"`
	// }
	gorm.Model
	UserID uint   `gorm:"not null;index"`
	Title  string `gorm:"not null"`
}

type GalleryService interface{}

type GalleryDB interface {
	Create(gallery *Gallery) error
}

type galleryGorm struct {
}
