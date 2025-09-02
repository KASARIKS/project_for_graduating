package validate_nextdate

import (
	"errors"
	"strconv"
	"strings"
)

func validateRepeatM(repeat string) error {
	splitedRepeatOnlySpaces := strings.Split(repeat, " ")
	if len(splitedRepeatOnlySpaces) < 2 {
		return errors.New("unsupported format length for m option")
	}

	// Check for days numbers
	days := strings.Split(splitedRepeatOnlySpaces[1], ",")
	if err := validateDaysNumbersForM(days); err != nil {
		return err
	}

	// Check for months numbers
	if len(splitedRepeatOnlySpaces) > 2 {
		months := strings.Split(splitedRepeatOnlySpaces[2], ",")
		if err := validateMonthNumbersForM(months, days); err != nil {
			return err
		}
	}

	return nil
}

func validateDaysNumbersForM(days []string) error {
	for _, v := range days {
		day, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("not numbers in days in repeat for m option")
		}
		if day > 31 || (day < 1 && day != -1 && day != -2 && day != -31) {
			return errors.New("too big or too small numbers in days in repeat for m option")
		}
	}

	return nil
}

func validateMonthNumbersForM(months, days []string) error {
	countErrsForMaxMonthDays := 0

	for _, v := range months {
		month, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("not numbers in months in repeat for m option")
		}
		if month > 12 || month < 1 {
			return errors.New("too big or too small numbers in months in repeat for m option")
		}

		// For different months different days quantity
		if err := validateDaysForMonth(month, days); err != nil {
			countErrsForMaxMonthDays++

			if countErrsForMaxMonthDays == len(months) {
				return err
			}
		}
	}

	return nil
}

func validateDaysForMonth(month int, days []string) error {
	for _, v := range days {
		day, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		if day > getMonthMaxQuantityDays(month) {
			return errors.New("too big value for month")
		}
	}

	return nil
}

func getMonthMaxQuantityDays(month int) int {
	months := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	return months[month-1]
}
