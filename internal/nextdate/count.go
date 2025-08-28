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
		dstartTime, err = countForWOption(dstartTime, splitedRepeat)
	case "m":
		dstartTime, err = countForMOption(dstartTime, repeat)
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

func countForWOption(dstartTime time.Time, splitedRepeat []string) (time.Time, error) {
	choosedWeekDays, err := getWeekDays(splitedRepeat[1:])
	if err != nil {
		return time.Time{}, err
	}

	dstartTime = dstartTime.AddDate(0, 0, 1)

	for !slices.Contains(choosedWeekDays, int(dstartTime.Weekday())) {
		dstartTime = dstartTime.AddDate(0, 0, 1)
	}

	return dstartTime, nil
}

func countForMOption(dstartTime time.Time, repeat string) (time.Time, error) {
	parts := strings.Split(repeat, " ")

	if len(parts) == 2 {
		days, err := getMonthDays(parts[1])
		if err != nil {
			return time.Time{}, err
		}

		for !slices.Contains(days, dstartTime.Day()) {
			dstartTime = dstartTime.AddDate(0, 0, 1)
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

		for !slices.Contains(months, int(dstartTime.Month())) || !slices.Contains(days, dstartTime.Day()) {
			dstartTime = dstartTime.AddDate(0, 0, 1)
		}
	}

	return dstartTime, nil
}
