package handler

import (
	"fmt"
	"strconv"
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
		startTimeStr := strconv.Itoa(modal.StartTime)
		endTimeStr := strconv.Itoa(modal.EndTime)

		dto := SlotSettingsDto{
			ID:        modal.ID,
			DayOfWeek: modal.DayOfWeek,
			StartTime: fmt.Sprintf("%s:%s", startTimeStr[:2], startTimeStr[2:]),
			EndTime:   fmt.Sprintf("%s:%s", endTimeStr[:2], endTimeStr[2:]),
		}

		//1235
		//

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
