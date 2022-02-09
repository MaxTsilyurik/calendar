package command

import (
	"calendar/internal/app/repository"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type DeleteEventHandler struct {
	repository repository.EventRepository
}

func NewDeleteEventHandler(repository repository.EventRepository) (DeleteEventHandler, error) {
	if repository == nil {
		panic("event repository is nil")
	}

	return DeleteEventHandler{repository: repository}, nil
}

func (h DeleteEventHandler) Handle(ctx context.Context, eventId uuid.UUID) (err error) {
	defer func() {
		err = errors.Wrapf(err, "delete event by id: %s", eventId.String())
	}()

	err = h.repository.DeleteBy(ctx, eventId)
	return
}
