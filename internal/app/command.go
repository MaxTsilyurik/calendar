package app

import (
	"calendar/internal/domain/calendar"
	"time"

	"github.com/google/uuid"
)

type (
	CreateEventCommand struct {
		Title            string
		TimeAndDateEvent time.Time
		EventDuration    time.Time
		Description      string
		UserId           uuid.UUID
		Reminder         calendar.ReminderType
	}

	EditEventCommand struct {
		Title            string
		TimeAndDateEvent time.Time
		EventDuration    time.Time
		Description      string
		Reminder         calendar.ReminderType
	}
)
