package main

import (
	"calendar/internal/app"
	"fmt"
	"github.com/pkg/errors"
)

func main() {

	err := errors.Wrapf(app.ErrNotFoundEvent, "getting event by %d", 10)

	if errors.Is(err, app.ErrNotFoundEvent) {
		fmt.Println("hello world")
	}
}
