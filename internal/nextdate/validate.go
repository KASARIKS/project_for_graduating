package nextdate

import (
	"errors"
	"strconv"
	"strings"
)

func ValidateRepeatParam(repeat string) error {
	splitedRepeat := SplitRepeat(repeat)
	if len(splitedRepeat) < 1 {
		return errors.New("empty repeat")
	}

	option := splitedRepeat[0]

	if err := validateRepeatbyOption(option, repeat); err != nil {
		return err
	}

	return nil
}

func validateRepeatbyOption(option string, repeat string) error {
	switch option {
	case "d":
		return validateRepeatD(repeat)
	case "y":
		return validateRepeatY(repeat)
	case "w":
		return validateRepeatW(repeat)
	case "m":
		return validateRepeatM(repeat)
	default:
		return errors.New("unsupported option")
	}
}

func validateRepeatD(repeat string) error {
	splitedRepeat := SplitRepeat(repeat)
	if len(splitedRepeat) != 2 {
		return errors.New("unsupported format length for d option")
	}

	return validateNumbersForD(splitedRepeat[1])
}

func validateNumbersForD(number string) error {
	n, err := strconv.Atoi(number)
	if err != nil {
		return errors.New("not numbers in repeat for d option")
	}
	if n > 400 {
		return errors.New("too big number for d option")
	}

	return nil
}

func validateRepeatY(repeat string) error {
	splitedRepeat := SplitRepeat(repeat)
	if len(splitedRepeat) > 1 {
		return errors.New("unsupported format length for y option")
	}

	return nil
}

func validateRepeatW(repeat string) error {
	splitedRepeat := SplitRepeat(repeat)
	if len(splitedRepeat) < 2 {
		return errors.New("unsupported format length for w option")
	}

	return validateNumbersForW(splitedRepeat[1:])
}

func validateNumbersForW(numbers []string) error {
	for _, v := range numbers {
		n, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("not numbers in repeat for w option")
		}
		if n > 7 || n < 1 {
			return errors.New("incorrect numbers in repeat for w option")
		}
	}

	return nil
}

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
