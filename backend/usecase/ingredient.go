package usecase

import (
	"context"
	"ref/entity"
)

type IngredientUseCase struct {
	IngredientRepogitory
}

func NewIngredientUseCase(repogitory IngredientRepogitory) *IngredientUseCase {
	return &IngredientUseCase{repogitory}
}

func (u IngredientUseCase) GetIngredients(ctx context.Context) ([]*entity.Ingredient, error) {
	return u.IngredientRepogitory.GetIngredients(ctx)
}
