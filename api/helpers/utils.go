package helpers

import "time"

// DateFormatter format date into the givin format
func DateFormatter(date string, format string) (time.Time, error) {
	t, err := time.Parse(format, date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
