package app

import (
	"time"

	"github.com/google/uuid"
)

type (
	DayEvents     time.Time
	WeekStartDay  time.Time
	MonthStartDay time.Time

	CommonEvent struct {
		Id        uuid.UUID
		Title     string
		TimeEvent time.Time
		UserId    uuid.UUID
	}
)
