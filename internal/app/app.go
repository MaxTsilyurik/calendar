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
		GetEventForDay   getEventForDate
		GetEventOfWeek   getEventForDate
		GetEventForMonth getEventForDate
	}

	getEventForDate interface {
		Handle(ctx context.Context, day time.Time) ([]*CommonEvent, error)
	}
)
