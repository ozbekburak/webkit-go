package converter

import (
	"log"
	"strconv"
	"time"
)

func Epoch(webkitTimestamp string) int64 {
	wt, err := strconv.ParseInt(webkitTimestamp, 10, 64)
	if err != nil {
		log.Printf("Failed to parse webkit timestamp: %s", err)
		return 0
	}

	return (wt / 1000000) - 11644473600
}

func Unix(webkitTimestamp string) time.Time {
	unixTime := time.Unix(Epoch(webkitTimestamp), 0)

	return unixTime
}

func ISO8601(webkitTimestamp string) string {
	return Unix(webkitTimestamp).Format("2006-01-02T15:04:05Z")
}
