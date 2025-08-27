package nextdate

import (
	"slices"
	"strconv"
	"strings"
)

func getWeekDays(splitedWeekDays []string) ([]int, error) {
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

func splitRepeat(repeat string) []string {
	firstSplit := strings.Split(repeat, " ")
	lastSplit := []string{}

	for _, v := range firstSplit {
		lastSplit = append(lastSplit, strings.Split(v, ",")...)
	}

	return lastSplit
}
