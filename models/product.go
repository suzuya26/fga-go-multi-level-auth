package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required~Your title is required"`
	Description string `json:"description" form:"description" valid:"required~Your description is required"`
	UserID      uint
	User        *User
}

// hooks
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return
}
