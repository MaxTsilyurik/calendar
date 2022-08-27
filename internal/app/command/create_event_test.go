package command

import (
	"calendar/internal/adapter/repository/mock"
	"calendar/internal/app"
	"calendar/internal/domain/calendar"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	testCases := []struct {
		Name          string
		Command       app.CreateEventCommand
		ErrorExpected error
	}{
		{
			Name: "create event",
			Command: app.CreateEventCommand{
				UserId:           uuid.New(),
				Title:            "new title",
				TimeAndDateEvent: time.Date(2022, time.April, 5, 12, 0, 0, 0, time.UTC),
				EventDuration:    time.Date(2022, time.April, 5, 12, 5, 0, 0, time.UTC),
				Description:      "test",
				Reminder:         calendar.ReminderType(2),
			},
			ErrorExpected: nil,
		},
		{
			Name: "dont create event when title not filled",
			Command: app.CreateEventCommand{
				UserId:           uuid.New(),
				TimeAndDateEvent: time.Date(2022, time.April, 5, 12, 0, 0, 0, time.UTC),
				EventDuration:    time.Date(2022, time.April, 5, 12, 5, 0, 0, time.UTC),
				Description:      "test",
				Reminder:         calendar.ReminderType(2),
			},
			ErrorExpected: calendar.ErrEventEmptyTitle,
		},
		{
			Name: "dont create event when reminder not valid",
			Command: app.CreateEventCommand{
				UserId:           uuid.New(),
				Title:            "new title",
				TimeAndDateEvent: time.Date(2022, time.April, 5, 12, 0, 0, 0, time.UTC),
				EventDuration:    time.Date(2022, time.April, 5, 12, 5, 0, 0, time.UTC),
				Description:      "test",
				Reminder:         0,
			},
			ErrorExpected: calendar.ErrEventNotValidReminderType,
		},
		{
			Name: "don't create event when end date and time before event date and time",
			Command: app.CreateEventCommand{
				UserId:           uuid.New(),
				Title:            "new title",
				TimeAndDateEvent: time.Date(2022, time.April, 5, 12, 0, 0, 0, time.UTC),
				EventDuration:    time.Date(2022, time.April, 5, 11, 59, 59, 0, time.UTC),
				Description:      "test",
				Reminder:         calendar.TimeOfEvent,
			},
			ErrorExpected: calendar.ErrEventNotValidEventDates,
		},
	}

	for i := range testCases {
		value := testCases[i]
		t.Run(value.Name, func(t *testing.T) {
			t.Parallel()

			handler := NewCreateEventHandler(mock.NewInMemoryEventRepository())
			eventId, err := handler.Handle(context.TODO(), value.Command)
			if err != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, value.ErrorExpected))
				return
			}

			require.NoError(t, err)
			require.NotEmpty(t, eventId)

		})
	}
}
