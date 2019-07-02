package models

import (
	"outstagram/server/dtos/dtomodels"
	"time"

	"github.com/jinzhu/gorm"
)

// User entity
type User struct {
	gorm.Model
	Username   string  `gorm:"unique;not null;unique_index"`
	Password   string  `gorm:"not null"`
	Fullname   string  `gorm:"not null"`
	Phone      *string `gorm:"unique"`
	Email      string  `gorm:"unique; not null"`
	AvatarURL  *string
	LastLogin  *time.Time
	Gender     bool
	NotifBoard NotifBoard `gorm:"association_autoupdate:false"`
	StoryBoard StoryBoard `gorm:"association_autoupdate:false"`
}

func (u *User) ToUserDTO() dtomodels.User {
	return dtomodels.User{
		ID:        u.ID,
		Fullname:  u.Fullname,
		Username:  u.Username,
		Gender:    u.Gender,
		Phone:     u.Phone,
		Email:     u.Email,
		AvatarURL: u.AvatarURL,
	}
}

func (u *User) ToMeDTO() dtomodels.Me {
	return dtomodels.Me{
		ID:        u.ID,
		Fullname:  u.Fullname,
		Username:  u.Username,
		Gender:    u.Gender,
		Phone:     u.Phone,
		Email:     u.Email,
		AvatarURL: u.AvatarURL,
	}
}

func (u *User) ToBasicUserDTO() dtomodels.BasicUser {
	return dtomodels.BasicUser{
		ID:       u.ID,
		Fullname: u.Fullname,
		Username: u.Username,
	}
}
