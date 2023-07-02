package repogitory

import (
	"context"
	"ref/entity"
	"ref/infrastructure/driver"
	"ref/infrastructure/repogitory/model"

	"gorm.io/gorm"
)

type ComponentRepogitory struct {
	*driver.TokenDriver
}

func NewComponentRepogitory(tokenDriver *driver.TokenDriver) *ComponentRepogitory {
	return &ComponentRepogitory{tokenDriver}
}

func (r ComponentRepogitory) GetComponents(ctx context.Context) ([]*entity.Component, error) {
	records := []*model.Component{}
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}
	var components []*entity.Component
	for _, record := range records {
		components = append(components, record.ToEntity())
	}
	return components, nil
}

func (r ComponentRepogitory) GetComponentByID(ctx context.Context, componentId string) (*entity.Component, error) {
	var record *model.Component
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.First(&record, "id = ?", componentId).Error; err != nil {
		return nil, err
	}
	return record.ToEntity(), nil
}
