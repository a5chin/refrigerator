package controller

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=./mock.go -package=$GOPACKAGE

import (
	"context"
	"ref/entity"
)

type IngredientUseCase interface {
	GetIngredients(ctx context.Context) ([]*entity.Ingredient, error)
}