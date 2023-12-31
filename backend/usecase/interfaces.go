package usecase

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock.go -package=$GOPACKAGE

import (
	"context"
	"ref/entity"
)

type IngredientRepository interface {
	GetIngredients(ctx context.Context, min, max *uint) ([]*entity.Ingredient, error)
	GetIngredientByID(ctx context.Context, ingredientId string) (*entity.Ingredient, error)
	UpdateIngredients(ctx context.Context, ingredientId string, weight uint) error
}

type NutritionRepository interface {
	GetNutritions(ctx context.Context) ([]*entity.Nutrition, error)
	GetNutritionByID(ctx context.Context, nutritionId string) (*entity.Nutrition, error)
}
