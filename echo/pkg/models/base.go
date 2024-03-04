package models

import "gorm.io/gorm"

type Base struct {
	gorm.Model
}

func NewBasemodel() Base {
	return Base{}
}
