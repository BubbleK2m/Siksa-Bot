package support

import "time"

func ParseTime(src string) (time.Time, error) {
	tme, err := time.Parse(time.RFC3339, src)
	return tme, err
}
