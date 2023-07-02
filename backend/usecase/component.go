package usecase

import (
	"context"
	"ref/entity"
)

type ComponentUseCase struct {
	ComponentRepogitory
}

func NewComponentUseCase(repogitory ComponentRepogitory) *ComponentUseCase {
	return &ComponentUseCase{repogitory}
}

func (u ComponentUseCase) GetComponents(ctx context.Context) ([]*entity.Component, error) {
	return u.ComponentRepogitory.GetComponents(ctx)
}
