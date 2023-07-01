package model

import (
	"ref/entity"

	"gorm.io/gorm"
)

type Ingredient struct {
	ID   string
	Name string
	gorm.Model
}

func (m Ingredient) ToEntity() *entity.Ingredient {
	return &entity.Ingredient{
		ID:   m.ID,
		Name: m.Name,
	}
}