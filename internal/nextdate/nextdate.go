package nextdate

import (
	"errors"
	"slices"
	"strconv"
	"time"
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	dstartTime, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}
	if err := validateRepeatParam(repeat); err != nil {
		return "", err
	}

	countedTime, err := countNextDate(now, dstartTime, repeat)
	if err != nil {
		return "", err
	}

	return countedTime.Format("20060102"), nil
}

func validateRepeatParam(repeat string) error {
	firstLetters := []byte{'d', 'y'}
	option := repeat[0]

	if !slices.Contains(firstLetters, option) {
		return errors.New("unsupported option")
	}

	if err := validateRepeatbyOption(option, repeat); err != nil {
		return err
	}

	return nil
}

func validateRepeatbyOption(option byte, repeat string) error {
	switch option {
	case 'd':
		if len(repeat) < 3 && repeat[1] != ' ' {
			return errors.New("unsupported format for d option")
		}
		n, err := strconv.Atoi(repeat[2:])
		if err != nil {
			return errors.New("unsupported format for d option")
		}
		if n > 400 {
			return errors.New("too big number for d option")
		}
	case 'y':
		if len(repeat) > 1 {
			return errors.New("unsupported format for y option")
		}
	}

	return nil
}

func countNextDate(now time.Time, dstartTime time.Time, repeat string) (time.Time, error) {
	option := repeat[0]
	switch option {
	case 'd':
		daysQuantity, err := strconv.Atoi(repeat[2:])
		if err != nil {
			return time.Time{}, err
		}

		for dstartTime.Before(now) {
			dstartTime = dstartTime.AddDate(0, 0, daysQuantity)
		}
	case 'y':
		for dstartTime.Before(now) {
			dstartTime = dstartTime.AddDate(1, 0, 0)
		}
	}

	return dstartTime, nil
}
