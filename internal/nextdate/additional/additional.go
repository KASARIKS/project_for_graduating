package additional_nextdate

import (
	"slices"
	"strconv"
	"strings"
)

func GetWeekDays(splitedWeekDays []string) ([]int, error) {
	var choosedWeekDays []int
	for _, v := range splitedWeekDays {
		day, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
		// Convert to Weekday numbers type
		day %= 7

		choosedWeekDays = append(choosedWeekDays, day)
	}
	slices.Sort(choosedWeekDays)

	return choosedWeekDays, nil
}

func SplitRepeat(repeat string) []string {
	firstSplit := strings.Split(repeat, " ")
	lastSplit := []string{}

	for _, v := range firstSplit {
		lastSplit = append(lastSplit, strings.Split(v, ",")...)
	}

	return lastSplit
}

func ConvertMonthDays(days string) ([]int, error) {
	splitedDays := strings.Split(days, ",")

	var daysInt []int
	for _, v := range splitedDays {
		day, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		daysInt = append(daysInt, day)
	}
	slices.Sort(daysInt)

	return daysInt, nil
}

func ConvertMonths(months string) ([]int, error) {
	splitedMonths := strings.Split(months, ",")

	var monthsInt []int
	for _, v := range splitedMonths {
		month, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		monthsInt = append(monthsInt, month)
	}
	slices.Sort(monthsInt)

	return monthsInt, nil
}
