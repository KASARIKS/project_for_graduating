package nextdate

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}
	if err := validateRepeatParam(repeat); err != nil {
		return "", err
	}

	countedTime, err := countNextDate(now, date, repeat)
	if err != nil {
		return "", err
	}

	return countedTime.Format(DateFormat), nil
}

func validateRepeatParam(repeat string) error {
	firstLetters := []string{"d", "y", "w"}
	splitedRepeat := splitRepeat(repeat)
	if len(splitedRepeat) < 1 {
		return errors.New("empty repeat")
	}

	option := splitedRepeat[0]

	if !slices.Contains(firstLetters, option) {
		return errors.New("unsupported option")
	}

	if err := validateRepeatbyOption(option, splitedRepeat); err != nil {
		return err
	}

	return nil
}

func validateRepeatbyOption(option string, splitedRepeat []string) error {
	switch option {
	case "d":
		if len(splitedRepeat) != 2 {
			return errors.New("unsupported format length for d option")
		}
		n, err := strconv.Atoi(splitedRepeat[1])
		if err != nil {
			return errors.New("not numbers in repeat for d option")
		}
		if n > 400 {
			return errors.New("too big number for d option")
		}
	case "y":
		if len(splitedRepeat) > 1 {
			return errors.New("unsupported format for y option")
		}
	case "w":
		if len(splitedRepeat) < 2 {
			return errors.New("unsupported format for d option")
		}
		// Check for numbers
		for _, v := range splitedRepeat[1:] {
			n, err := strconv.Atoi(v)
			if err != nil {
				return errors.New("not numbers in repeat for w option")
			}
			if n > 7 || n < 1 {
				return errors.New("incorrect numbers in repeat for w option")
			}
		}
	}

	return nil
}

func countNextDate(now time.Time, dstartTime time.Time, repeat string) (time.Time, error) {
	splitedRepeat := splitRepeat(repeat)
	option := splitedRepeat[0]

	// Made described algorithm for d and y option, but it works not like you want.
	// Maybe algorithm was described incorrect.
	switch option {
	case "d":
		daysQuantity, err := strconv.Atoi(splitedRepeat[1])
		if err != nil {
			return time.Time{}, err
		}

		for !dstartTime.After(now) {
			dstartTime = dstartTime.AddDate(0, 0, daysQuantity)
		}
	case "y":
		for !dstartTime.After(now) {
			dstartTime = dstartTime.AddDate(1, 0, 0)
		}
	case "w":
		choosedWeekDays, err := getWeekDays(splitedRepeat[1:])
		if err != nil {
			return time.Time{}, err
		}

		dstartTime = dstartTime.AddDate(0, 0, 1)

		for !slices.Contains(choosedWeekDays, int(dstartTime.Weekday())) {
			dstartTime = dstartTime.AddDate(0, 0, 1)
		}
	}

	return dstartTime, nil
}

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
