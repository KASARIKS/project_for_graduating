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
	if option == "m" {

	}

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
		splitedRepeat := splitRepeat(repeat)
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
		splitedRepeat := splitRepeat(repeat)
		if len(splitedRepeat) > 1 {
			return errors.New("unsupported format length for y option")
		}
	case "w":
		splitedRepeat := splitRepeat(repeat)
		if len(splitedRepeat) < 2 {
			return errors.New("unsupported format length for w option")
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
	case "m":
		splitedRepeatOnlySpaces := strings.Split(repeat, " ")
		if len(splitedRepeatOnlySpaces) < 2 {
			return errors.New("unsupported format length for m option")
		}

		// Check for days numbers
		days := strings.Split(splitedRepeatOnlySpaces[1], ",")
		for _, v := range days {
			day, err := strconv.Atoi(v)
			if err != nil {
				return errors.New("not numbers in days in repeat for m option")
			}
			if day > 31 || day < -31 {
				return errors.New("too big or too small numbers in days in repeat for m option")
			}
		}

		// Check for months numbers
		if len(splitedRepeatOnlySpaces) > 2 {
			months := strings.Split(splitedRepeatOnlySpaces[2], ",")
			for _, v := range months {
				month, err := strconv.Atoi(v)
				if err != nil {
					return errors.New("not numbers in months in repeat for m option")
				}
				if month > 12 || month < 1 {
					return errors.New("too big or too small numbers in months in repeat for m option")
				}
			}
		}
	}

	return nil
}
