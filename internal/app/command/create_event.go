package command

import (
	"calendar/internal/app"
	"calendar/internal/app/repository"
	"calendar/internal/domain/calendar"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CreateEventHandler struct {
	eventRepository repository.EventRepository
}

func NewCreateEventHandler(repository repository.EventRepository) CreateEventHandler {
	if repository == nil {
		panic("event repository is nil")
	}

	return CreateEventHandler{eventRepository: repository}
}

func (h CreateEventHandler) Handle(ctx context.Context, cec app.CreateEventCommand) (eventID string, err error) {

	defer func() {
		err = errors.Wrapf(err, "event creation by user: %s", cec.UserId.String())
	}()

	eventID = uuid.NewString()

	event, err := calendar.NewEvent(eventID, cec.UserId.String(),
		cec.Title, cec.Description, cec.TimeAndDateEvent, cec.EventDuration, cec.Reminder)

	if err != nil {
		return "", err
	}

	if err := h.eventRepository.Save(ctx, event); err != nil {
		return "", err
	}

	return

}
