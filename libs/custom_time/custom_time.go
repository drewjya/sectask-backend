package custom_time

import "time"

func ParseFromUTC(date string) string {
	t, _ := time.Parse(time.RFC3339, date)
	return t.Add(7 * time.Hour).Format(time.RFC3339)
}
