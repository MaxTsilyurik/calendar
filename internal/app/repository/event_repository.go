package repository

import (
	"calendar/internal/domain/calendar"
	"context"
	"time"
)

type EventRepository interface {
	Save(ctx context.Context, event *calendar.Event) error
	Update(ctx context.Context, event *calendar.Event) error
	DeleteBy(ctx context.Context, eventId string) error
	GetEventByDay(ctx context.Context, date time.Time) ([]*calendar.Event, error)
	GetEventByWeekStart(ctx context.Context, dateWeek time.Time) ([]*calendar.Event, error)
	GetEventByMonthStart(ctx context.Context, dateMonth time.Time) ([]*calendar.Event, error)
	FindById(ctx context.Context, eventId string) (*calendar.Event, error)
}
