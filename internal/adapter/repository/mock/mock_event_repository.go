package mock

import (
	"calendar/internal/domain/calendar"
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type MockRepository struct {
	db map[string]*calendar.Event
	mu *sync.RWMutex
}

func NewMockRepository() MockRepository {
	db := make(map[string]*calendar.Event)
	return MockRepository{db: db}
}

func (m MockRepository) Save(ctx context.Context, event *calendar.Event) error {
	m.mu.Lock()
	m.db[event.Id()] = event
	m.mu.Unlock()
	return nil
}

func (m MockRepository) Update(ctx context.Context, event *calendar.Event) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.db[event.Id()]
	if !ok {
		return errors.New("not found event")
	}
	m.db[event.Id()] = event
	return nil
}

func (m MockRepository) DeleteBy(ctx context.Context, eventId uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.db[eventId.String()]
	if !ok {
		return errors.New("not found event")
	}
	delete(m.db, eventId.String())
	return nil

}

func (m MockRepository) GetEventByDay(ctx context.Context, date time.Time) ([]*calendar.Event, error) {

	m.mu.RLock()
	defer m.mu.RUnlock()
	events := make([]*calendar.Event, 0, len(m.db))
	for _, v := range m.db {
		if compareDate(date, v.TimeAndDateEvent()) {
			events = append(events, v)
		}
	}
	if len(events) == 0 {
		return nil, errors.Errorf("Event not found by date: %s", date.Format(calendar.LayoutDateISO))
	}
	return events, nil
}

func (m MockRepository) GetEventByWeekStart(ctx context.Context, dateWeek time.Time) ([]calendar.Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return nil, nil
}

func (m MockRepository) GetEventByMonthStart(ctx context.Context, dateMonth time.Time) ([]calendar.Event, error) {
	return nil, nil
}

func (m MockRepository) FindById(ctx context.Context, eventId uuid.UUID) (*calendar.Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	event, ok := m.db[eventId.String()]
	if !ok {
		return nil, errors.Errorf("not found event by id: %s", eventId.String())
	}
	return event, nil
}

func compareDate(dateOne, dateTwo time.Time) bool {
	y1, m1, d1 := dateOne.Date()
	y2, m2, d2 := dateTwo.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
