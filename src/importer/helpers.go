package importer

import "time"

func cleanupDescription(s string) string {
	return s
}

func parseDate(s string) (time.Time, error) {
	return time.Parse("01/02/2006", s)
}
