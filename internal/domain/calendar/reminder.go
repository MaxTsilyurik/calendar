package calendar

type ReminderType uint8

const (
	TimeOfEvent ReminderType = iota + 1
	InFiveMinutes
	InTenMinutes
	ThirtyMinutes
	InAnHour
	InTwoHours
	OneDayBefore
)

func (t ReminderType) isValid() bool {
	switch t {
	case TimeOfEvent, InFiveMinutes, InTenMinutes, ThirtyMinutes, InAnHour, InTwoHours, OneDayBefore:
		return true
	default:
		return false
	}
}
