package main

import (
	"fmt"
	"time"
)

func main() {
	// currentTime := time.Now()
	// fmt.Println(currentTime)
	// fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05 PM"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000 PM"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000 PM"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000000 PM"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000000 PM -0700"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000000 PM -07:00"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000000 PM -0700 MST"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000000 PM -07:00 MST"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000000 PM -0700 MST Mon"))

	// fmt.Println(currentTime.Year())
	// fmt.Println(currentTime.Month())
	// fmt.Println(currentTime.Day())
	// fmt.Println(currentTime.Hour())
	// fmt.Println(currentTime.Minute())
	// fmt.Println(currentTime.Second())
	// fmt.Println(currentTime.Nanosecond())
	// fmt.Println(currentTime.Weekday())
	// fmt.Println(currentTime.Location())
	// fmt.Println(currentTime.Zone())

	// // UTC = Coordinated Universal Time
	// fmt.Println(currentTime.UTC())

	// // Local = Local Time Zone
	// fmt.Println(currentTime.Local())

	// // Add(d Duration) Time
	// fmt.Println(currentTime.Add(24 * time.Hour))

	// // AddDate(years int, months int, days int) Time
	// fmt.Println(currentTime.AddDate(1, 0, 0))

	// // Sub(time Time) Duration
	// fmt.Println(currentTime.Sub(currentTime.Add(24 * time.Hour)))

	// // Before(t Time) bool
	// fmt.Println(currentTime.Before(currentTime.Add(24 * time.Hour)))
	// fmt.Println(currentTime.After(currentTime.Add(24 * time.Hour)))
	// fmt.Println(currentTime.Equal(currentTime.Add(24 * time.Hour)))
	// fmt.Println(currentTime.Round(time.Hour))

	// // Truncate(d Duration) Time
	// fmt.Println(currentTime.Truncate(time.Hour))

	// // YearDay() int
	// fmt.Println(currentTime.YearDay())
	// fmt.Println(currentTime.ISOWeek())
	// fmt.Println(currentTime.Unix())
	// fmt.Println(currentTime.UnixNano())

	// fmt.Println(currentTime.Format(time.RFC3339))
	// fmt.Println(currentTime.Format(time.RFC3339Nano))
	// fmt.Println(currentTime.Format(time.RFC822))
	// fmt.Println(currentTime.Format(time.RFC822Z))
	// fmt.Println(currentTime.Format(time.RFC850))
	// fmt.Println(currentTime.Format(time.RFC1123))
	// fmt.Println(currentTime.Format(time.RFC1123Z))
	// fmt.Println(currentTime.Format(time.ANSIC))
	// fmt.Println(currentTime.Format(time.UnixDate))
	// fmt.Println(currentTime.Format(time.RubyDate))
	// fmt.Println(currentTime.Format(time.Kitchen))
	// fmt.Println(currentTime.Format(time.Stamp))
	// fmt.Println(currentTime.Format(time.StampMilli))
	// fmt.Println(currentTime.Format(time.StampMicro))
	// fmt.Println(currentTime.Format(time.StampNano))

	// fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05 PM"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000 PM"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000 PM"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000000 PM"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000000 PM -0700"))
	// fmt.Println(currentTime.Format("2006-01-02 03:04:05.000000000 PM -07:00"))

	// Parse(layout, value string) (Time, error)

	// timeString := "2021-08-01 15:04:05"
	// ts, err := time.Parse("2006-01-02 15:04:05", timeString)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(ts)
	// }

	// Monotonic Clock

	t1 := time.Now()
	fmt.Println(t1)
	time.Sleep(time.Second)
	t2 := time.Now()
	fmt.Println(time.Now())
	fmt.Println(t2.Sub(t1))

	// The above code measures the time taken to execute the code between t1 and t2
}
