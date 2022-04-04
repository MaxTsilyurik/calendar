package calendar

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewEvent(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		Name   string
		Params struct {
			id               string
			title            string
			timeAndDateEvent time.Time
			eventDuration    time.Time
			description      string
			userId           string
			reminder         ReminderType
		}
		ExpectedErr error
	}{
		{
			Name: "valid_create_event",
			Params: struct {
				id               string
				title            string
				timeAndDateEvent time.Time
				eventDuration    time.Time
				description      string
				userId           string
				reminder         ReminderType
			}{
				id:               uuid.NewString(),
				title:            "New event",
				timeAndDateEvent: time.Date(2022, 02, 05, 22, 33, 12, 0, time.UTC),
				eventDuration:    time.Date(2022, 02, 05, 23, 33, 12, 0, time.UTC),
				description:      "Description event",
				userId:           uuid.NewString(),
				reminder:         InFiveMinutes,
			},
			ExpectedErr: nil,
		},
		{
			Name: "empty_id_event",
			Params: struct {
				id               string
				title            string
				timeAndDateEvent time.Time
				eventDuration    time.Time
				description      string
				userId           string
				reminder         ReminderType
			}{
				title:            "New event",
				timeAndDateEvent: time.Date(2022, 02, 05, 22, 33, 12, 0, time.UTC),
				eventDuration:    time.Date(2022, 02, 05, 23, 33, 12, 0, time.UTC),
				description:      "Description event",
				userId:           uuid.NewString(),
				reminder:         InFiveMinutes,
			},
			ExpectedErr: ErrEventEmptyId,
		},
		{
			Name: "empty_userId_event",
			Params: struct {
				id               string
				title            string
				timeAndDateEvent time.Time
				eventDuration    time.Time
				description      string
				userId           string
				reminder         ReminderType
			}{
				id:               uuid.NewString(),
				title:            "New event",
				timeAndDateEvent: time.Date(2022, 02, 05, 22, 33, 12, 0, time.UTC),
				eventDuration:    time.Date(2022, 02, 05, 23, 33, 12, 0, time.UTC),
				description:      "Description event",
				reminder:         InFiveMinutes,
			},
			ExpectedErr: ErrEventEmptyUserId,
		},
		{
			Name: "empty_title_event",
			Params: struct {
				id               string
				title            string
				timeAndDateEvent time.Time
				eventDuration    time.Time
				description      string
				userId           string
				reminder         ReminderType
			}{
				id:               uuid.NewString(),
				timeAndDateEvent: time.Date(2022, 02, 05, 22, 33, 12, 0, time.UTC),
				eventDuration:    time.Date(2022, 02, 05, 23, 33, 12, 0, time.UTC),
				description:      "Description event",
				userId:           uuid.NewString(),
				reminder:         InFiveMinutes,
			},
			ExpectedErr: ErrEventEmptyTitle,
		},
		{
			Name: "not_valid_reminderType_event",
			Params: struct {
				id               string
				title            string
				timeAndDateEvent time.Time
				eventDuration    time.Time
				description      string
				userId           string
				reminder         ReminderType
			}{
				id:               uuid.NewString(),
				title:            "New event",
				timeAndDateEvent: time.Date(2022, 02, 05, 22, 33, 12, 0, time.UTC),
				eventDuration:    time.Date(2022, 02, 05, 23, 33, 12, 0, time.UTC),
				description:      "Description event",
				userId:           uuid.NewString(),
				reminder:         ReminderType(100),
			},
			ExpectedErr: ErrEventNotValidReminderType,
		},
		{

			Name: "eventDuration_equals_timeAndDateEvent",
			Params: struct {
				id               string
				title            string
				timeAndDateEvent time.Time
				eventDuration    time.Time
				description      string
				userId           string
				reminder         ReminderType
			}{
				id:               uuid.NewString(),
				title:            "New event",
				timeAndDateEvent: time.Date(2022, 02, 05, 22, 33, 12, 0, time.UTC),
				eventDuration:    time.Date(2022, 02, 05, 22, 33, 12, 0, time.UTC),
				description:      "Description event",
				userId:           uuid.NewString(),
				reminder:         InFiveMinutes,
			},
			ExpectedErr: nil,
		},
		{

			Name: "eventDuration_before_timeAndDateEvent",
			Params: struct {
				id               string
				title            string
				timeAndDateEvent time.Time
				eventDuration    time.Time
				description      string
				userId           string
				reminder         ReminderType
			}{
				id:               uuid.NewString(),
				title:            "New event",
				timeAndDateEvent: time.Date(2022, 02, 05, 22, 33, 12, 0, time.UTC),
				eventDuration:    time.Date(2022, 02, 04, 23, 33, 12, 0, time.UTC),
				description:      "Description event",
				userId:           uuid.NewString(),
				reminder:         InFiveMinutes,
			},
			ExpectedErr: ErrEventNotValidEventDates,
		},
	}

	for index := range testCase {
		value := testCase[index]
		t.Run(value.Name, func(t *testing.T) {
			t.Parallel()

			event, err := NewEvent(value.Params.id, value.Params.userId,
				value.Params.title, value.Params.description, value.Params.timeAndDateEvent,
				value.Params.eventDuration, value.Params.reminder)

			if err != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, value.ExpectedErr))
				return
			}
			require.NoError(t, err)
			require.Equal(t, event.id, value.Params.id)
			require.Equal(t, event.userId, value.Params.userId)
			require.Equal(t, event.title, value.Params.title)
			require.Equal(t, event.description, value.Params.description)
			require.Equal(t, event.eventDuration, value.Params.eventDuration)
			require.Equal(t, event.timeAndDateEvent, value.Params.timeAndDateEvent)
			require.Equal(t, event.reminder, value.Params.reminder)
		})
	}

}
