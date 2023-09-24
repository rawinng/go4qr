package utilities

import (
	"time"
)

func FormatDefaultTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
