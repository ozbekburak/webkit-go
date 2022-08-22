package converter

import (
	"log"
	"strconv"
	"time"
)

func Unix(webkitTimestamp string) time.Time {
	wt, err := strconv.ParseInt(webkitTimestamp, 10, 64)
	if err != nil {
		log.Printf("Failed to parse webkit timestamp: %s", err)
		return time.Time{}
	}

	unixTime := time.Unix((wt/1000000)-11644473600, 0)

	return unixTime
}
