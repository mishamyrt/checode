package dates

import "time"

const SHORT_DATE_LAYOUT = "2006-01-02"
const GB_DATE_LAYOUT = "1 January 2006"

// ParseShort date format
func ParseShort(date string) (time.Time, error) {
	return time.Parse(SHORT_DATE_LAYOUT, date)
}

// Format date
func Format(date time.Time) string {
	return date.Format(GB_DATE_LAYOUT)
}
