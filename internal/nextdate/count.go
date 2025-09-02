package nextdate

import (
	"slices"
	"strconv"
	"strings"
	"time"
)

func countNextDate(now time.Time, dstartTime time.Time, repeat string) (time.Time, error) {
	splitedRepeat := splitRepeat(repeat)
	option := splitedRepeat[0]
	var err error

	// Made described algorithm for d and y option, but it works not like you want.
	// Maybe algorithm was described incorrect.
	// Don't understand working with dstartTime, and don't understand why I need 2 time parameters.
	switch option {
	case "d":
		dstartTime, err = countForDOption(now, dstartTime, splitedRepeat)
	case "y":
		dstartTime, err = countForYOption(now, dstartTime)
	case "w":
		dstartTime, err = countForWOption(now, dstartTime, splitedRepeat)
	case "m":
		dstartTime, err = countForMOption(now, dstartTime, repeat)
	}

	return dstartTime, err
}

func countForDOption(now time.Time, dstartTime time.Time, splitedRepeat []string) (time.Time, error) {
	daysQuantity, err := strconv.Atoi(splitedRepeat[1])
	if err != nil {
		return time.Time{}, err
	}

	dstartTime = dstartTime.AddDate(0, 0, daysQuantity)

	for !dstartTime.After(now) {
		dstartTime = dstartTime.AddDate(0, 0, daysQuantity)
	}

	return dstartTime, nil
}

func countForYOption(now time.Time, dstartTime time.Time) (time.Time, error) {
	dstartTime = dstartTime.AddDate(1, 0, 0)

	for !dstartTime.After(now) {
		dstartTime = dstartTime.AddDate(1, 0, 0)
	}

	return dstartTime, nil
}

func countForWOption(now time.Time, dstartTime time.Time, splitedRepeat []string) (time.Time, error) {
	choosedWeekDays, err := getWeekDays(splitedRepeat[1:])
	if err != nil {
		return time.Time{}, err
	}

	dstartTime = dstartTime.AddDate(0, 0, 1)
	for !dstartTime.After(now) {
		dstartTime = dstartTime.AddDate(0, 0, 1)
	}

	for !slices.Contains(choosedWeekDays, int(dstartTime.Weekday())) {
		dstartTime = dstartTime.AddDate(0, 0, 1)
	}

	return dstartTime, nil
}

func countForMOption(now time.Time, dstartTime time.Time, repeat string) (time.Time, error) {
	parts := strings.Split(repeat, " ")

	if len(parts) == 2 {
		days, err := getMonthDays(parts[1])
		if err != nil {
			return time.Time{}, err
		}

		dstartTime = dstartTime.AddDate(0, 0, 1)
		for !dstartTime.After(now) {
			dstartTime = dstartTime.AddDate(0, 0, 1)
		}

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
		months, err := getMonths(parts[2])
		if err != nil {
			return time.Time{}, err
		}

		days, err := getMonthDays(parts[1])
		if err != nil {
			return time.Time{}, err
		}

		dstartTime = dstartTime.AddDate(0, 0, 1)
		for !dstartTime.After(now) {
			dstartTime = dstartTime.AddDate(0, 0, 1)
		}

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
