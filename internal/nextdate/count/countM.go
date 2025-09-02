package count_nextdate

import (
	"slices"
	"strings"
	"time"

	additional_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/additional"
)

func CountForMOption(now time.Time, dstartTime time.Time, repeat string) (time.Time, error) {
	parts := strings.Split(repeat, " ")

	if len(parts) == 2 {
		days, err := additional_nextdate.ConvertMonthDays(parts[1])
		if err != nil {
			return time.Time{}, err
		}

		dstartTime = addDays(now, dstartTime, 1)

		minusDays := getMinusDays(days)
		for !slices.Contains(days, dstartTime.Day()) {
			dstartTime = dstartTime.AddDate(0, 0, 1)

			t := time.Date(dstartTime.Year(), dstartTime.Month(), 32, 0, 0, 0, 0, time.UTC)
			daysInMonth := 32 - t.Day()

			reverseDays := getDaysByMinusDays(minusDays, daysInMonth)
			if slices.Contains(reverseDays, dstartTime.Day()) {
				break
			}
		}
	} else if len(parts) == 3 {
		months, err := additional_nextdate.ConvertMonths(parts[2])
		if err != nil {
			return time.Time{}, err
		}

		days, err := additional_nextdate.ConvertMonthDays(parts[1])
		if err != nil {
			return time.Time{}, err
		}

		dstartTime = addDays(now, dstartTime, 1)

		for !slices.Contains(months, int(dstartTime.Month())) || !slices.Contains(days, dstartTime.Day()) {
			dstartTime = dstartTime.AddDate(0, 0, 1)
		}
	}

	return dstartTime, nil
}

func getMinusDays(days []int) []int {
	minusDays := []int{}
	for _, v := range days {
		if v < 0 {
			minusDays = append(minusDays, v)
		}
	}

	return minusDays
}

func getDaysByMinusDays(minusDays []int, daysInMonth int) []int {
	days := []int{}
	for _, v := range minusDays {
		days = append(days, daysInMonth+(v+1))
	}

	return days
}
