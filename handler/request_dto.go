package handler

import (
	"time"

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

func toBookingModel(dto CreateBookingRequestDTO, userID int64) model.Booking {
	bookingDate, _ := time.Parse("01/02/2003", dto.BookingDate)
	return model.Booking{
		UserID:      userID,
		CalendarID:  dto.CalendarID,
		Status:      "pending",
		BookingDate: bookingDate,
	}
}
