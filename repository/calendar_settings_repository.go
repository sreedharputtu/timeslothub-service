package repository

import (
	"github.com/sreedharputtu/timeslothub-service/model"
	"gorm.io/gorm"
)

type CalendarSettingsRepository interface {
	Save(model.CalendarSettings) (int64, error)
	FindByID(calendarID int64) (model.CalendarSettings, error)
	FindByUserID(userID int64) ([]model.CalendarSettings, error)
}

type CalendarSettingsRepositoryImpl struct {
	DB *gorm.DB
}

func NewCalendarRepoImpl(DB *gorm.DB) CalendarSettingsRepository {
	return &CalendarSettingsRepositoryImpl{DB: DB}
}

// FindByUserID implements CalendarSettingsRepository.
func (cri *CalendarSettingsRepositoryImpl) FindByUserID(userID int64) ([]model.CalendarSettings, error) {
	var calendars []model.CalendarSettings
	result := cri.DB.Where("user_id=?", userID).Find(&calendars)
	if result.Error != nil {
		return calendars, result.Error
	}
	return calendars, nil
}

// FindByID implements CalendarSettingsRepository.
func (cri *CalendarSettingsRepositoryImpl) FindByID(calendarID int64) (model.CalendarSettings, error) {
	var calendar model.CalendarSettings
	result := cri.DB.Where("id=?", calendarID).Find(&calendar)
	if result.Error != nil {
		return calendar, result.Error
	}
	return calendar, nil
}

// Save implements CalendarSettingsRepository.
func (cri *CalendarSettingsRepositoryImpl) Save(cal model.CalendarSettings) (int64, error) {
	result := cri.DB.Save(&cal)
	return cal.ID, result.Error
}
