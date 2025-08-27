package nextdate

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

func validateRepeatParam(repeat string) error {
	firstLetters := []string{"d", "y", "w", "m"}
	splitedRepeat := splitRepeat(repeat)
	if len(splitedRepeat) < 1 {
		return errors.New("empty repeat")
	}

	option := splitedRepeat[0]

	if !slices.Contains(firstLetters, option) {
		return errors.New("unsupported option")
	}

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
	splitedRepeat := splitRepeat(repeat)
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
	splitedRepeat := splitRepeat(repeat)
	if len(splitedRepeat) > 1 {
		return errors.New("unsupported format length for y option")
	}

	return nil
}

func validateRepeatW(repeat string) error {
	splitedRepeat := splitRepeat(repeat)
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
		if err := validateMonthNumbersForM(months); err != nil {
			return err
		}
	}

	return nil
}

func validateDaysNumbersForM(numbers []string) error {
	for _, v := range numbers {
		day, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("not numbers in days in repeat for m option")
		}
		if day > 31 || day < -31 {
			return errors.New("too big or too small numbers in days in repeat for m option")
		}
	}

	return nil
}

func validateMonthNumbersForM(numbers []string) error {
	for _, v := range numbers {
		month, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("not numbers in months in repeat for m option")
		}
		if month > 12 || month < 1 {
			return errors.New("too big or too small numbers in months in repeat for m option")
		}
	}

	return nil
}
