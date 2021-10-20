package dates

import "time"

const ShortDateLayout = "2006-01-02"
const CLIDateLayout = "1 January 2006"

// ParseShort date format.
func ParseShort(date string) (time.Time, error) {
	return time.Parse(ShortDateLayout, date)
}

// Format date.
func Format(date time.Time) string {
	return date.Format(CLIDateLayout)
}
