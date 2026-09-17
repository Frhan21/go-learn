package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"

    schedule, err := time.Parse(layout, date)
    if err != nil {
        return time.Time{}
    }
    return schedule
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

    schedule, err := time.Parse(layout, date)
    if err != nil {
        return false
    }

    return time.Now().After(schedule)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	schedule, err := time.Parse(layout, date) 
    if err != nil {
        return false
    }

    hourSchedule := schedule.Hour()

    return hourSchedule >= 12 && hourSchedule < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    schedule, err := time.Parse(layout, date)
    if err != nil {
        return ""
    }

    outputLayout := "Monday, January 2, 2006, at 15:04."
	return fmt.Sprintf("You have an appointment on %s", schedule.Format(outputLayout))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(2026, time.September, 15, 0, 0, 0, 0, time.UTC)
}
