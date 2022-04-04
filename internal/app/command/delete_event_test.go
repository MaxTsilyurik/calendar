package command

import (
	"calendar/internal/adapter/repository/mock"
	"calendar/internal/app"
	"calendar/internal/domain/calendar"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDeleteEventHandel(t *testing.T) {
	t.Parallel()

	eventId := uuid.New()
	event, _ := calendar.NewEvent(eventId.String(), uuid.NewString(), "new string", "new string",
		time.Now(), time.Now(), calendar.TimeOfEvent)

	testCases := []struct {
		Name          string
		EventSave     *calendar.Event
		EventId       uuid.UUID
		ErrorExpected error
	}{
		{
			Name:          "don't delete when event not found",
			EventId:       uuid.New(),
			EventSave:     nil,
			ErrorExpected: app.ErrEventNotFound,
		},
		{
			Name:          "should delete event",
			EventId:       eventId,
			EventSave:     event,
			ErrorExpected: app.ErrEventNotFound,
		},
	}

	for i := range testCases {
		value := testCases[i]
		t.Run(value.Name, func(t *testing.T) {
			t.Parallel()

			repository := mock.NewInMemoryEventRepository()
			handler := NewDeleteEventHandler(repository)

			if value.EventSave != nil {
				_ = repository.Save(context.TODO(), value.EventSave)
			}

			err := handler.Handle(context.TODO(), value.EventId)
			if err != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, value.ErrorExpected))
				return
			}

			require.NoError(t, err)
		})
	}
}
