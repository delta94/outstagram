package models

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/enums/postprivacy"
)

// Post entity
type Post struct {
	gorm.Model
	Content       *string
	Privacy       postPrivacy.Privacy `gorm:"default:0"`
	CommentableID uint
	ReactableID   uint
	ViewableID    uint
	UserID        uint
	User          User
	ImageID       uint `gorm:"not null"`
	Image         Image
	Images        []PostImage `gorm:"foreignkey:PostID"`
	ImageCount    int         `gorm:"-"`
	Popularity    float32     `gorm:"popularity;default:0	"`
}
