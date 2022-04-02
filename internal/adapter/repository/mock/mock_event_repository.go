package mock

import (
	"calendar/internal/app"
	"calendar/internal/domain/calendar"
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type InMemoryEventRepository struct {
	db map[string]*calendar.Event
	mu *sync.RWMutex
}

func NewInMemoryEventRepository() *InMemoryEventRepository {
	db := make(map[string]*calendar.Event)
	return &InMemoryEventRepository{db: db, mu: &sync.RWMutex{}}
}

func (m *InMemoryEventRepository) Save(ctx context.Context, event *calendar.Event) error {
	m.mu.Lock()
	m.db[event.Id()] = event
	m.mu.Unlock()
	return nil
}

func (m *InMemoryEventRepository) Update(ctx context.Context, event *calendar.Event) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.db[event.Id()]
	if !ok {
		return errors.Wrapf(app.ErrEventNotFound, "update event by id: %s", event.Id())
	}
	m.db[event.Id()] = event
	return nil
}

func (m *InMemoryEventRepository) DeleteBy(ctx context.Context, eventId string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.db[eventId]
	if !ok {
		return errors.Wrapf(app.ErrEventNotFound, "removing event by id: %s", eventId)
	}
	delete(m.db, eventId)
	return nil

}

func (m *InMemoryEventRepository) GetEventByDay(ctx context.Context, date time.Time) ([]*calendar.Event, error) {

	m.mu.RLock()
	defer m.mu.RUnlock()
	events := make([]*calendar.Event, 0, len(m.db))
	for _, v := range m.db {
		if compareDate(date, v.TimeAndDateEvent()) {
			events = append(events, v)
		}
	}
	if len(events) == 0 {
		return nil, errors.Wrapf(app.ErrEventNotFound, "getting event by date: %s", date.Format(calendar.LayoutDateISO))
	}
	return events, nil
}

func (m *InMemoryEventRepository) GetEventByWeekStart(ctx context.Context, dateWeek time.Time) ([]*calendar.Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	endDate := dateWeek.AddDate(0, 0, 6)
	events, err := m.betweenDate(dateWeek, endDate)
	if err != nil {
		return nil, errors.Wrapf(err,
			"getting an event from the start of the week between a date :%s  %s",
			dateWeek.Format(calendar.LayoutDateISO), endDate.Format(calendar.LayoutDateISO),
		)
	}
	return events, nil
}

func (m *InMemoryEventRepository) betweenDate(dateStart time.Time, dateEnd time.Time) ([]*calendar.Event, error) {
	events := make([]*calendar.Event, 0, len(m.db))
	for _, event := range m.db {
		if dateInBetween(dateStart, dateEnd, event.TimeAndDateEvent()) {
			events = append(events, event)
		}
	}
	if len(events) == 0 {
		return nil, app.ErrEventNotFound
	}
	return events, nil
}

func (m *InMemoryEventRepository) GetEventByMonthStart(ctx context.Context, dateMonth time.Time) ([]*calendar.Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	endDate := dateMonth.AddDate(0, 1, -1)
	events, err := m.betweenDate(dateMonth, endDate)
	if err != nil {
		return nil, errors.Wrapf(err,
			"getting an event from the start of the month between a date :%s  %s",
			dateMonth.Format(calendar.LayoutDateISO), endDate.Format(calendar.LayoutDateISO),
		)
	}
	return events, nil
}

func (m *InMemoryEventRepository) FindById(ctx context.Context, eventId string) (*calendar.Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	event, ok := m.db[eventId]
	if !ok {
		return nil, errors.Wrapf(app.ErrEventNotFound, "getting event by %s", eventId)
	}
	return event, nil
}

func compareDate(dateOne, dateTwo time.Time) bool {
	y1, m1, d1 := dateOne.Date()
	y2, m2, d2 := dateTwo.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func dateInBetween(dateStart, dateEnd, compareDate time.Time) bool {
	return dateStart == compareDate || dateEnd == compareDate ||
		(dateStart.Before(compareDate) && dateEnd.After(compareDate))
}
