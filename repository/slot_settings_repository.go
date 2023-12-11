package repository

import (
	"github.com/sreedharputtu/timeslothub-service/model"
	"gorm.io/gorm"
)

type SlotSettingsRepository interface {
	Save(slotSettings model.SlotSettings) error
}

type SlotSettingsRepositoryImpl struct {
	Db *gorm.DB
}

func NewSlotSettingsRepository(Db *gorm.DB) SlotSettingsRepository {
	return &SlotSettingsRepositoryImpl{Db: Db}
}

func (ssr *SlotSettingsRepositoryImpl) Save(slotSettings model.SlotSettings) error {
	result := ssr.Db.Create(&slotSettings)
	return result.Error
}

func (ur *SlotSettingsRepositoryImpl) FindAll() ([]model.SlotSettings, error) {
	var slotSettings []model.SlotSettings
	result := ur.Db.Find(&slotSettings)
	return slotSettings, result.Error
}
