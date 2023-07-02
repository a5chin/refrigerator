package controller

import (
	"net/http"
	"ref/entity"

	"github.com/gin-gonic/gin"
)

type ComponentController struct {
	ComponentUseCase
}

func NewComponentController(u ComponentUseCase) *ComponentController {
	return &ComponentController{u}
}

type GetComponentsResponse struct {
	Components []*entity.Component `json:"components"`
}

// GetComponents godoc
//
// @Summary	全栄養素取得 API
// @Description
// @Tags		Component
// @Accept		json
// @Produce		json
// @Success		200		{object}	GetComponentsResponse	"OK"
// @Failure		401		{object}	entity.ErrorResponse
// @Failure		404		{object}	entity.ErrorResponse
// @Router		/components	[get]
func (c ComponentController) GetComponents(ctx *gin.Context) (interface{}, error) {
	components, err := c.ComponentUseCase.GetComponents(ctx)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return GetComponentsResponse{Components: components}, nil
}

type GetComponentResponse struct {
	Component *entity.Component `json:"component"`
}

// GetComponentByID godoc
//
// @Summary	栄養素取得 API
// @Description
// @Tags		Component
// @Accept		json
// @Param		componentId	path		string	true			"栄養素 ID"
// @Produce		json
// @Success		200			{object}	GetComponentsResponse	"OK"
// @Failure		401			{object}	entity.ErrorResponse
// @Failure		404			{object}	entity.ErrorResponse
// @Router		/components/{componentId}	[get]
func (c ComponentController) GetComponentByID(ctx *gin.Context) (interface{}, error) {
	componentId := ctx.Param("componentId")
	component, err := c.ComponentUseCase.GetComponentByID(ctx, componentId)
	if err != nil {
		return nil, entity.WrapError(http.StatusBadRequest, err)
	}
	return GetComponentResponse{Component: component}, nil
}
