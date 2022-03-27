package mock

import (
	"calendar/internal/app"
	"calendar/internal/domain/calendar"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestEventInMockRepository(t *testing.T) {
	t.Parallel()
	repository := NewInMemoryEventRepository()

	t.Run("save new event", func(t *testing.T) {
		t.Parallel()
		event, _ := calendar.NewEvent(uuid.NewString(), uuid.NewString(), "new string", "new string",
			time.Now(), time.Now(), calendar.TimeOfEvent)
		err := repository.Save(context.TODO(), event)
		require.Nil(t, err)
	})

	t.Run("find event", func(t *testing.T) {
		t.Parallel()
		eventId := uuid.NewString()

		event, _ := calendar.NewEvent(eventId, uuid.NewString(), "new string", "new string",
			time.Now(), time.Now(), calendar.TimeOfEvent)

		_ = repository.Save(context.TODO(), event)

		testCases := []struct {
			Name  string
			Id    string
			Err   error
			Event *calendar.Event
		}{
			{
				Name:  "should find event by id",
				Id:    eventId,
				Event: event,
				Err:   nil,
			},
			{
				Name:  "not should find event by id",
				Id:    uuid.NewString(),
				Event: nil,
				Err:   app.ErrNotFoundEvent,
			},
		}

		for i := range testCases {
			testCase := testCases[i]

			t.Run(testCase.Name, func(t *testing.T) {
				t.Parallel()
				event, err := repository.FindById(context.TODO(), testCase.Id)
				if err != nil {
					require.Error(t, err)
					require.True(t, errors.Is(err, testCase.Err))
					return
				}

				require.Equal(t, testCase.Event.Id(), event.Id())
				require.Equal(t, testCase.Event.CreatedUser(), event.CreatedUser())
				require.Equal(t, testCase.Event.EventDuration(), event.EventDuration())
				require.Equal(t, testCase.Event.TimeAndDateEvent(), event.TimeAndDateEvent())
				require.Equal(t, testCase.Event.Title(), event.Title())
				require.Equal(t, testCase.Event.Description(), event.Description())
				require.Equal(t, testCase.Event.Reminder(), event.Reminder())

			})
		}

	})

	t.Run("update event", func(t *testing.T) {
		testCases := []struct {
			Name string
		}{}
	})
}
