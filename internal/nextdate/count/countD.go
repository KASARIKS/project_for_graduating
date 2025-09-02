package count_nextdate

import (
	"strconv"
	"time"
)

func CountForDOption(now, dstartTime time.Time, splitedRepeat []string) (time.Time, error) {
	daysQuantity, err := strconv.Atoi(splitedRepeat[1])
	if err != nil {
		return time.Time{}, err
	}

	dstartTime = addDays(now, dstartTime, daysQuantity)

	return dstartTime, nil
}
