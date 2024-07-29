package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"not null" json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url"`
	UserID   uint64 `json:"user_id"`
	User     *User
	GormModel
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
