package repogitory

import (
	"context"
	"fmt"
	"log"
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

func (r IngredientRepogitory) GetIngredients(ctx context.Context, min, max *uint) ([]*entity.Ingredient, error) {
	records := []*model.Ingredient{}
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	log.Print("aslkjgasdldgjasdga", *min, max)
	if min != nil {
		db = db.Where(fmt.Sprintf("weight >= %d", *min))
	}
	if max != nil {
		log.Print("lkadsgalksjdgagjldasjg")
		db = db.Where(fmt.Sprintf("weight <= %d", *max))
	}
	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}
	var ingredients []*entity.Ingredient
	for _, record := range records {
		ingredients = append(ingredients, record.ToEntity())
	}
	return ingredients, nil
}
