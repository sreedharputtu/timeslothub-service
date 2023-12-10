package repository

import "github.com/sreedharputtu/timeslothub-service/model"

type SlotSettingsRepository interface {
	Save(slotSettings model.SlotSettings)
}

type SlotSettingsRepositoryImpl struct {
}
