package utilities

import (
	"time"
)

// Get Time Now
func DateTimeNow() time.Time {

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	return time.Now().UTC().Local()

}

// Add hours
func DateTimeNowAddHoursUnix(hours int64) int64 {

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	return time.Now().Add(time.Hour * time.Duration(hours)).Unix()

}

// Add hours
func DateTimeNowAddHours(hours int64) time.Time {

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	return time.Now().Add(time.Hour * time.Duration(hours))

}
