package query

import (
	"calendar/internal/app"
	"calendar/internal/app/repository"
	"context"
	"time"
)

type EventForMonthHandler struct {
	repository repository.EventRepository
}

func NewEventForMonthHandler(eventRepository repository.EventRepository) EventForDayHandler {
	if eventRepository == nil {
		panic("event repository is nil")
	}
	return EventForDayHandler{
		repository: eventRepository,
	}
}

func (h EventForMonthHandler) Handle(ctx context.Context, day time.Time) ([]*app.CommonEvent, error) {
	eventByDay, err := h.repository.GetEventByMonthStart(ctx, day)
	if err != nil {
		return []*app.CommonEvent{}, err
	}
	return app.UnmarshalCommonEvents(eventByDay)
}
