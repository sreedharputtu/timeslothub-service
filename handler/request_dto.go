package handler

import (
	"time"

	"github.com/labstack/gommon/log"
	"github.com/sreedharputtu/timeslothub-service/model"
)

type createUserRequestDto struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type SlotSettingsDto struct {
	ID        int64  `json:"id"`
	DayOfWeek string `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	UserID    int64  `json:"user_id"`
}

func convertSlotSettings(slotSettingsList []model.SlotSettings) []SlotSettingsDto {
	slotSettingsDtoList := make([]SlotSettingsDto, len(slotSettingsList))
	for i, modal := range slotSettingsList {
		startTimeStr := modal.StartTime.Format("15:04")
		endTimeStr := modal.EndTime.Format("15:04")
		dto := SlotSettingsDto{
			ID:        modal.ID,
			DayOfWeek: modal.DayOfWeek,
			StartTime: startTimeStr,
			EndTime:   endTimeStr,
		}
		slotSettingsDtoList[i] = dto
	}
	return slotSettingsDtoList
}

type CreateBookingRequestDTO struct {
	CalendarID  int64  `json:"calendar_id"`
	BookingDate string `json:"booking_date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Email       string `json:"email"`
}

func toBookingModel(dto CreateBookingRequestDTO, userID int64, createdBy int64) model.Booking {
	bookingDate, err := time.Parse("2006-01-02", dto.BookingDate)
	if err != nil {
		log.Error(err)
	}
	startTime, err := time.Parse("15:04", dto.StartTime)
	if err != nil {
		log.Error(err)
	}
	endTime, err := time.Parse("15:04", dto.EndTime)
	if err != nil {
		log.Error(err)
	}
	return model.Booking{
		UserID:        userID,
		CalendarID:    dto.CalendarID,
		Status:        "pending",
		StartDateTime: startTime,
		EndDateTime:   endTime,
		BookingDate:   bookingDate,
		CreatedBy:     createdBy,
	}
}
