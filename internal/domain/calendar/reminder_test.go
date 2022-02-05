package calendar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		Name     string
		Reminder ReminderType
		IsValid  bool
	}{
		{
			Name:     "valid_reminder_type",
			Reminder: InAnHour,
			IsValid:  true,
		},
		{
			Name:     "in_valid_reminder_type",
			Reminder: ReminderType(100),
			IsValid:  false,
		},
	}

	for index := range testCase {
		value := testCase[index]
		t.Run(value.Name, func(t *testing.T) {
			t.Parallel()

			ok := value.Reminder.isValid()

			require.Equal(t, ok, value.IsValid)
		})
	}

}
