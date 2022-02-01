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

type GalleryService interface {
	GalleryDB
}

type GalleryDB interface {
	Create(gallery *Gallery) error
}

func NewGalleryService(db *gorm.DB) GalleryService {
	return &galleryService{
		GalleryDB: &galleryValidator{&galleryGorm{db}},
	}
}

type galleryService struct {
	GalleryDB
}

type galleryValidator struct {
	GalleryDB
}

var _ GalleryDB = &galleryGorm{}

type galleryGorm struct {
	db *gorm.DB
}

func (gg *galleryGorm) Create(gallery *Gallery) error {
	return gg.db.Create(gallery).Error
}

// Update will update provided gallery with all of the data
// in the provided gallery object.
//func (gg *userGorm) Update(gallery Gallery) error {
//	return gg.db.Save(gallery).Error
//}

// Delete will delete the gallery with the provided id
//func (gg *userGorm) Delete(id uint) error {
//	gallery := User{Model: gorm.Model{ID: id}}
//	return gg.db.Delete(&gallery).Error
//}
