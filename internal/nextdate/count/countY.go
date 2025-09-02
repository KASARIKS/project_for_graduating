package count_nextdate

import "time"

func CountForYOption(now, dstartTime time.Time) (time.Time, error) {
	dstartTime = addYear(now, dstartTime)

	return dstartTime, nil
}

func addYear(now, dstartTime time.Time) time.Time {
	dstartTime = dstartTime.AddDate(1, 0, 0)

	for !dstartTime.After(now) {
		dstartTime = dstartTime.AddDate(1, 0, 0)
	}

	return dstartTime
}
