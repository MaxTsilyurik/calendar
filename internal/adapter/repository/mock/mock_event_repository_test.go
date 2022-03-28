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
				Err:   app.ErrEventNotFound,
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

		eventId := uuid.NewString()
		userIdOld := uuid.NewString()
		userIdNew := uuid.NewString()

		eventUpdate, _ := calendar.NewEvent(eventId, userIdOld, "new string", "new string",
			time.Now(), time.Now(), calendar.TimeOfEvent)

		eventUpdateCaseOne, _ := calendar.NewEvent(eventId, userIdNew, "new string update", "new string update",
			time.Now(), time.Now(), calendar.TimeOfEvent)

		eventUpdateCaseTwo, _ := calendar.NewEvent(uuid.NewString(), userIdNew, "new string update", "new string update",
			time.Now(), time.Now(), calendar.TimeOfEvent)

		testCases := []struct {
			Name        string
			Event       *calendar.Event
			EventUpdate *calendar.Event
			Err         error
		}{
			{
				Name:        "update should",
				Event:       eventUpdate,
				EventUpdate: eventUpdateCaseOne,
				Err:         nil,
			},
			{
				Name:        "not found event",
				Event:       nil,
				EventUpdate: eventUpdateCaseTwo,
				Err:         app.ErrEventNotFound,
			},
		}

		for i := range testCases {
			testCase := testCases[i]
			t.Run(testCase.Name, func(t *testing.T) {
				t.Parallel()
				if testCase.Event != nil {
					_ = repository.Save(context.TODO(), testCase.Event)
				}

				err := repository.Update(context.TODO(), testCase.EventUpdate)
				if err != nil {
					require.Error(t, err)
					require.True(t, errors.Is(err, testCase.Err))
					return
				}

				event, _ := repository.FindById(context.TODO(), testCase.Event.Id())

				require.Equal(t, testCase.EventUpdate.Id(), event.Id())
				require.Equal(t, testCase.EventUpdate.CreatedUser(), event.CreatedUser())
				require.Equal(t, testCase.EventUpdate.EventDuration(), event.EventDuration())
				require.Equal(t, testCase.EventUpdate.TimeAndDateEvent(), event.TimeAndDateEvent())
				require.Equal(t, testCase.EventUpdate.Title(), event.Title())
				require.Equal(t, testCase.EventUpdate.Description(), event.Description())
				require.Equal(t, testCase.EventUpdate.Reminder(), event.Reminder())

			})
		}
	})

	t.Run("getting event by day", func(t *testing.T) {
		t.Parallel()
		date := time.Date(2022, time.March, 22, 15, 00, 46, 0, time.UTC)
		dateEnd := date.Add(5 * time.Minute)

		eventOne, _ := calendar.NewEvent(uuid.NewString(), uuid.NewString(), "new string", "new string",
			date, dateEnd, calendar.TimeOfEvent)

		eventTwo, _ := calendar.NewEvent(uuid.NewString(), uuid.NewString(), "new string update", "new string update",
			date, dateEnd, calendar.TimeOfEvent)

		eventThree, _ := calendar.NewEvent(uuid.NewString(), uuid.NewString(), "new string update", "new string update",
			date.AddDate(0, 1, 2), dateEnd.AddDate(0, 1, 2), calendar.TimeOfEvent)

		eventFour, _ := calendar.NewEvent(uuid.NewString(), uuid.NewString(), "new string update", "new string update",
			date.AddDate(0, 0, 2), dateEnd.AddDate(0, 0, 2), calendar.TimeOfEvent)

		listSave := []*calendar.Event{eventOne, eventTwo, eventThree, eventFour}

		findEvent := map[string]*calendar.Event{
			eventOne.Id(): eventOne,
			eventTwo.Id(): eventTwo,
		}

		testCases := []struct {
			Name           string
			DateSearch     time.Time
			CountFindEvent int
			FindEvent      map[string]*calendar.Event
			Err            error
		}{
			{
				Name:           "not found",
				DateSearch:     date.AddDate(0, 0, 10),
				CountFindEvent: 0,
				FindEvent:      nil,
				Err:            app.ErrEventNotFound,
			},
			{
				Name:           "should find events by date",
				DateSearch:     date,
				CountFindEvent: 2,
				FindEvent:      findEvent,
				Err:            nil,
			},
		}

		for _, event := range listSave {
			_ = repository.Save(context.TODO(), event)
		}

		for i := range testCases {
			testCase := testCases[i]
			t.Run(testCase.Name, func(t *testing.T) {
				t.Parallel()
				events, err := repository.GetEventByDay(context.TODO(), testCase.DateSearch)

				if err != nil {
					require.Error(t, err)
					require.True(t, errors.Is(err, testCase.Err))
					return
				}
				require.Equal(t, testCase.CountFindEvent, len(events))

				for _, event := range events {
					if value, ok := testCase.FindEvent[event.Id()]; ok {
						require.Equal(t, value.Id(), event.Id())
						require.Equal(t, value.CreatedUser(), event.CreatedUser())
						require.Equal(t, value.EventDuration(), event.EventDuration())
						require.Equal(t, value.TimeAndDateEvent(), event.TimeAndDateEvent())
						require.Equal(t, value.Title(), event.Title())
						require.Equal(t, value.Description(), event.Description())
						require.Equal(t, value.Reminder(), event.Reminder())
					}
				}

			})
		}
	})
}
