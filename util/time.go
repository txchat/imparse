package util

import "time"

var (
	shanghai = Location("Asia/Shanghai")  // Shanghai *time.Location
	hongkong = Location("Asia/Hong_Kong") // Hong Kong *time.Location
	local    = Location("Local")          // Local *time.Location
	utc      = Location("UTC")            // UTC *time.Location
)

// Location returns *time.Location by location name.
func Location(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return loc
}

// Shanghai returns Shanghai *time.Location.
func Shanghai() *time.Location {
	return shanghai
}

// TimeNowUnixNano returns now unix nanosecond timestamp.
func TimeNowUnixNano(location ...*time.Location) int64 {
	loc := Shanghai()
	if len(location) != 0 {
		loc = location[0]
	}
	return time.Now().In(loc).UnixNano()
}
