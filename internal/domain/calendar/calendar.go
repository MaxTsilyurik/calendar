package calendar

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

const (
	LayoutISO     = "2006-01-02 15:04:05"
	LayoutDateISO = "2006-01-02"
)

var (
	ErrEventEmptyId              = errors.New("id field not filled")
	ErrEventEmptyUserId          = errors.New("userId field not filled")
	ErrEventEmptyTitle           = errors.New("title field not filled")
	ErrEventNotValidReminderType = errors.New("reminder field is not valid")
	ErrEventNotValidEventDates   = errors.New("the end date and time of the event must be later than the start of the event")
)

type Event struct {
	id               string
	title            string
	timeAndDateEvent time.Time
	eventDuration    time.Time
	description      string
	userId           string
	reminder         ReminderType
}

func NewEvent(id, userId, title, description string, timeAndDateEvent, eventDuration time.Time, reminder ReminderType) (*Event, error) {

	if id == "" {
		return nil, ErrEventEmptyId
	}

	if userId == "" {
		return nil, ErrEventEmptyUserId
	}

	if title == "" {
		return nil, ErrEventEmptyTitle
	}

	if !reminder.isValid() {
		return nil, ErrEventNotValidReminderType
	}

	if eventDuration.Equal(timeAndDateEvent) || eventDuration.Before(timeAndDateEvent) {
		return nil, ErrEventNotValidEventDates
	}

	return &Event{
		id:               id,
		userId:           userId,
		title:            title,
		description:      description,
		timeAndDateEvent: timeAndDateEvent,
		eventDuration:    eventDuration,
		reminder:         reminder,
	}, nil
}

func (m *Event) String() string {
	return fmt.Sprintf("Event:{\n\tid: %v,\n\ttitle: %v,\n\ttimeAndDateEvent: %v,\n\teventDuration: %v,\n\tdescription: %v,\n\tuserId: %v,\n\treminder: %v\n}",
		m.id, m.title, m.timeAndDateEvent, m.eventDuration, m.description, m.userId, m.reminder)
}

func (m *Event) Id() string {
	return m.id
}

func (m *Event) Title() string {
	return m.title
}

func (m *Event) TimeAndDateEvent() time.Time {
	return m.timeAndDateEvent
}

func (m *Event) EventDuration() time.Time {
	return m.eventDuration
}

func (m *Event) Description() string {
	return m.description
}

func (m *Event) CreatedUser() string {
	return m.userId
}

func (m *Event) Reminder() ReminderType {
	return m.reminder
}
