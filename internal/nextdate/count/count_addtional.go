package count_nextdate

import "time"

func addDays(now, dstartTime time.Time, daysQuantity int) time.Time {
	dstartTime = dstartTime.AddDate(0, 0, daysQuantity)

	for !dstartTime.After(now) {
		dstartTime = dstartTime.AddDate(0, 0, daysQuantity)
	}

	return dstartTime
}
