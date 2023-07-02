package model

import (
	"ref/entity"
)

type Component struct {
	ID   string
	Name string
}

func (m Component) ToEntity() *entity.Component {
	return &entity.Component{
		ID:   m.ID,
		Name: m.Name,
	}
}
