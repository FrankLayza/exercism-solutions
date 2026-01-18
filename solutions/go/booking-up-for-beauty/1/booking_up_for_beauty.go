package booking

import (
    "time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    t, err:= time.Parse(layout,date)
    if(err != nil){
        fmt.Println(err)
        return time.Time{}
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t,err := time.Parse("January 2, 2006 15:04:05", date)
    if(err != nil){
        fmt.Println(err)
        return false
    }
    var now = time.Now()
    return (now.After(t))
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if(err != nil){
        fmt.Println(err)
        return false
    }
    hour:= t.Hour()
    return (hour >= 12 && hour < 18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
   t, err := time.Parse("1/2/2006 15:04:05", date)
   if err != nil {
    fmt.Println(err)
    return ""
   }
   hour := t.Hour()
   weekday := t.Weekday().String()
   month := t.Month().String()
   day := t.Day()
   year := t.Year()
   return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%02d.",
    weekday, month, day, year, hour, t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    anniversary := time.Date(time.Now().UTC().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
    return anniversary
}
