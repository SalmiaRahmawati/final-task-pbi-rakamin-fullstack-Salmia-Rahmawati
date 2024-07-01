package app

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormApp
	Title    string `gorm:"not null" json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url"`
	UserID   uint64 `json:"user_id"`
	User     *User
}

type PhotoInput struct {
	Title    string `json:"title" valid:"required"`
	Caption  string `json:"caption"`
	UserID   uint64 `json:"user_id" valid:"required"`
	PhotoURL string `json:"photo_url"`
}

type PhotoOutput struct {
	GormApp
	Title    string `json:"title" valid:"required"`
	Caption  string `json:"caption"`
	UserID   uint64 `json:"user_id" valid:"required"`
	PhotoURL string `json:"photo_url"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
