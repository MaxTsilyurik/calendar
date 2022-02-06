package app

import (
	"time"

	"github.com/google/uuid"
)

type (
	CommonEvent struct {
		Id        uuid.UUID
		Title     string
		TimeEvent time.Time
		UserId    uuid.UUID
	}
)
