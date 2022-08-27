package app

import (
	"calendar/internal/domain/calendar"
	"time"

	"github.com/google/uuid"
)

type (
	CreateEventCommand struct {
		Title            string                `json:"title,omitempty"`
		TimeAndDateEvent time.Time             `json:"time_and_date_event"`
		EventDuration    time.Time             `json:"event_duration"`
		Description      string                `json:"description,omitempty"`
		UserId           uuid.UUID             `json:"user_id,omitempty"`
		Reminder         calendar.ReminderType `json:"reminder,omitempty"`
	}

	EditEventCommand struct {
		Title            string
		TimeAndDateEvent time.Time
		EventDuration    time.Time
		Description      string
		Reminder         calendar.ReminderType
	}
)
