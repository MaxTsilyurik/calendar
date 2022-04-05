package app

import (
	"calendar/internal/domain/calendar"
	"github.com/pkg/errors"
	"time"

	"github.com/google/uuid"
)

type (
	CommonEvent struct {
		Id        uuid.UUID
		Title     string
		TimeEvent time.Time
		UserId    uuid.UUID
	}
)

func UnmarshalCommonEvents(events []*calendar.Event) ([]*CommonEvent, error) {
	eventsCommon := make([]*CommonEvent, len(events))
	for _, event := range events {
		commonEvent, err := unmarshallCommonEvent(event)
		if err != nil {
			return []*CommonEvent{}, nil
		}
		eventsCommon = append(eventsCommon, commonEvent)
	}
	return eventsCommon, nil
}

func unmarshallCommonEvent(event *calendar.Event) (*CommonEvent, error) {
	eventId, err := uuid.Parse(event.Id())
	if err != nil {
		return nil, errors.Wrapf(err, "papse exception event_id by event: %s", event.Id())
	}
	userId, err := uuid.Parse(event.CreatedUser())
	if err != nil {
		return nil, errors.Wrapf(err, "papse exception user_id by event: %s", event.CreatedUser())
	}
	return &CommonEvent{
		Id:        eventId,
		Title:     event.Title(),
		TimeEvent: event.TimeAndDateEvent(),
		UserId:    userId,
	}, nil
}
