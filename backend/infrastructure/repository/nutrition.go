package repository

import (
	"context"
	"ref/entity"
	"ref/infrastructure/driver"
	"ref/infrastructure/repository/model"

	"gorm.io/gorm"
)

type NutritionRepository struct {
	*driver.TokenDriver
}

func NewNutritionRepository(tokenDriver *driver.TokenDriver) *NutritionRepository {
	return &NutritionRepository{tokenDriver}
}

func (r NutritionRepository) GetNutritions(ctx context.Context) ([]*entity.Nutrition, error) {
	records := []*model.Nutrition{}
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}
	var nutritions []*entity.Nutrition
	for _, record := range records {
		nutritions = append(nutritions, record.ToEntity())
	}
	return nutritions, nil
}

func (r NutritionRepository) GetNutritionByID(ctx context.Context, nutritionId string) (*entity.Nutrition, error) {
	var record *model.Nutrition
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.First(&record, "id = ?", nutritionId).Error; err != nil {
		return nil, err
	}
	return record.ToEntity(), nil
}
