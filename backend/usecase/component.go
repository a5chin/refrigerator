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

func (u ComponentUseCase) GetComponentByID(ctx context.Context, componentId string) (*entity.Component, error) {
	return u.ComponentRepogitory.GetComponentByID(ctx, componentId)
}
