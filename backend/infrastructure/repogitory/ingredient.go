package repogitory

import (
	"context"
	"ref/entity"
	"ref/infrastructure/driver"
	"ref/infrastructure/repogitory/model"

	"gorm.io/gorm"
)

type IngredientRepogitory struct {
	*driver.TokenDriver
}

func NewIngredientRepogitory(tokenDriver *driver.TokenDriver) *IngredientRepogitory {
	return &IngredientRepogitory{tokenDriver}
}

func (r IngredientRepogitory) GetIngredients(ctx context.Context) ([]*entity.Ingredient, error) {
	records := []*model.Ingredient{}
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}
	var ingredients []*entity.Ingredient
	for _, record := range records {
		ingredients = append(ingredients, record.ToEntity())
	}
	return ingredients, nil
}
