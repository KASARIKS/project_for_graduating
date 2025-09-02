package count_nextdate

import (
	"slices"
	"time"

	additional_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/additional"
)

func CountForWOption(now time.Time, dstartTime time.Time, splitedRepeat []string) (time.Time, error) {
	choosedWeekDays, err := additional_nextdate.GetWeekDays(splitedRepeat[1:])
	if err != nil {
		return time.Time{}, err
	}

	dstartTime = addDays(now, dstartTime, 1)
	for !slices.Contains(choosedWeekDays, int(dstartTime.Weekday())) {
		dstartTime = dstartTime.AddDate(0, 0, 1)
	}

	return dstartTime, nil
}
