package repository

import (
	"github.com/sreedharputtu/timeslothub-service/model"
	"gorm.io/gorm"
)

type SlotSettingsRepository interface {
	Save(slotSettings model.SlotSettings) error
	FindByCalendarID(calendarID int64) ([]model.SlotSettings, error)
	FindAll() ([]model.SlotSettings, error)
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

func (ssr *SlotSettingsRepositoryImpl) FindAll() ([]model.SlotSettings, error) {
	var slotSettings []model.SlotSettings
	result := ssr.Db.Find(&slotSettings)
	return slotSettings, result.Error
}

func (ssr *SlotSettingsRepositoryImpl) FindByCalendarID(calendarID int64) ([]model.SlotSettings, error) {
	var slotSettingsList []model.SlotSettings
	err := ssr.Db.Where("calendar_id=?", calendarID).Find(&slotSettingsList).Error
	if err != nil {
		return slotSettingsList, err
	}
	return slotSettingsList, nil
}
