package controller

import (
	"net/http"
	"ref/entity"

	"github.com/gin-gonic/gin"
)

type IngredientController struct {
	IngredientUseCase
}

func NewIngredientController(u IngredientUseCase) *IngredientController {
	return &IngredientController{u}
}

type GetIngredientsResponse struct {
	Ingredients []*entity.Ingredient `json:"ingredients"`
}

// GetIngredients godoc
//
// @Summary	全素材取得 API
// @Description
// @Tags		Ingredient
// @Accept		json
// @Produce		json
// @Success		200		{object}	GetIngredientsResponse	"OK"
// @Failure		401		{object}	entity.ErrorResponse
// @Failure		404		{object}	entity.ErrorResponse
// @Router		/ingredients	[get]
func (c IngredientController) GetIngredients(ctx *gin.Context) (interface{}, error) {
	ingredients, err := c.IngredientUseCase.GetIngredients(ctx)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return GetIngredientsResponse{Ingredients: ingredients}, nil
}
