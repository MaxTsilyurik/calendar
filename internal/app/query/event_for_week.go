package query

import (
	"calendar/internal/app"
	"calendar/internal/app/repository"
	"context"
	"time"
)

type EventForWeekHandler struct {
	repository repository.EventRepository
}

func NewEventForWeekHandler(eventRepository repository.EventRepository) EventForDayHandler {
	if eventRepository == nil {
		panic("event repository is nil")
	}
	return EventForDayHandler{
		repository: eventRepository,
	}
}

func (h EventForWeekHandler) Handler(ctx context.Context, day time.Time) ([]*app.CommonEvent, error) {
	eventByDay, err := h.repository.GetEventByWeekStart(ctx, day)
	if err != nil {
		return []*app.CommonEvent{}, err
	}
	return app.UnmarshalCommonEvents(eventByDay)
}
