package nextdate

import (
	"slices"
	"strconv"
	"strings"
	"time"
)

func CountNextDate(now time.Time, dstartTime time.Time, repeat string) (time.Time, error) {
	splitedRepeat := SplitRepeat(repeat)
	option := splitedRepeat[0]
	var err error

	switch option {
	case "d":
		dstartTime, err = CountForDOption(now, dstartTime, splitedRepeat)
	case "y":
		dstartTime, err = CountForYOption(now, dstartTime)
	case "w":
		dstartTime, err = CountForWOption(now, dstartTime, splitedRepeat)
	case "m":
		dstartTime, err = CountForMOption(now, dstartTime, repeat)
	}

	return dstartTime, err
}

func addDays(now, dstartTime time.Time, daysQuantity int) time.Time {
	dstartTime = dstartTime.AddDate(0, 0, daysQuantity)

	for !dstartTime.After(now) {
		dstartTime = dstartTime.AddDate(0, 0, daysQuantity)
	}

	return dstartTime
}

func CountForDOption(now, dstartTime time.Time, splitedRepeat []string) (time.Time, error) {
	daysQuantity, err := strconv.Atoi(splitedRepeat[1])
	if err != nil {
		return time.Time{}, err
	}

	dstartTime = addDays(now, dstartTime, daysQuantity)

	return dstartTime, nil
}

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

func CountForWOption(now time.Time, dstartTime time.Time, splitedRepeat []string) (time.Time, error) {
	choosedWeekDays, err := GetWeekDays(splitedRepeat[1:])
	if err != nil {
		return time.Time{}, err
	}

	dstartTime = addDays(now, dstartTime, 1)
	for !slices.Contains(choosedWeekDays, int(dstartTime.Weekday())) {
		dstartTime = dstartTime.AddDate(0, 0, 1)
	}

	return dstartTime, nil
}

func CountForMOption(now time.Time, dstartTime time.Time, repeat string) (time.Time, error) {
	parts := strings.Split(repeat, " ")

	if len(parts) == 2 {
		days, err := ConvertMonthDays(parts[1])
		if err != nil {
			return time.Time{}, err
		}

		dstartTime = addDays(now, dstartTime, 1)

		dstartTime = theClosestDstart(days, nil, dstartTime)
	} else if len(parts) == 3 {
		months, err := ConvertMonths(parts[2])
		if err != nil {
			return time.Time{}, err
		}

		days, err := ConvertMonthDays(parts[1])
		if err != nil {
			return time.Time{}, err
		}

		dstartTime = addDays(now, dstartTime, 1)

		dstartTime = theClosestDstart(days, months, dstartTime)
	}

	return dstartTime, nil
}

func theClosestDstart(days, months []int, dstartTime time.Time) time.Time {
	minusDays := getMinusDays(days)

	for monthContain(months, dstartTime) || !slices.Contains(days, dstartTime.Day()) {
		dstartTime = dstartTime.AddDate(0, 0, 1)

		t := time.Date(dstartTime.Year(), dstartTime.Month(), 32, 0, 0, 0, 0, time.UTC)
		daysInMonth := 32 - t.Day()

		reverseDays := getDaysByMinusDays(minusDays, daysInMonth)
		if slices.Contains(reverseDays, dstartTime.Day()) {
			break
		}
	}

	return dstartTime
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

func monthContain(months []int, dstartTime time.Time) bool {
	if months == nil {
		return false
	}

	return !slices.Contains(months, int(dstartTime.Month()))
}

func getDaysByMinusDays(minusDays []int, daysInMonth int) []int {
	days := []int{}
	for _, v := range minusDays {
		days = append(days, daysInMonth+(v+1))
	}

	return days
}
