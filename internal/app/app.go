package app

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type (
	Commands struct {
		CreateEvent createEventHandler
		EditEvent   editEventHandler
		RemoveEvent removeEventHandler
	}

	//createEventHandler
	createEventHandler interface {
		Handle(ctx context.Context, cec CreateEventCommand) (string, error)
	}

	editEventHandler interface {
		Handle(ctx context.Context, eventId uuid.UUID, eec EditEventCommand) error
	}

	removeEventHandler interface {
		Handle(ctx context.Context, eventId uuid.UUID) error
	}
)

type (
	Queries struct {
		GetEventForDay   getEventForDayHandler
		GetEventOfWeek   getEventsOfWeekHandler
		GetEventForMonth getEventsForMonthHandler
	}

	getEventForDayHandler interface {
		Handle(ctx context.Context, day time.Time) []CommonEvent
	}

	getEventsOfWeekHandler interface {
		Handle(ctx context.Context, wsd time.Time) []CommonEvent
	}

	getEventsForMonthHandler interface {
		Handle(ctx context.Context, msd time.Time) []CommonEvent
	}
)
