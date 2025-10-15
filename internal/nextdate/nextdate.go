package nextdate

import (
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}

	if err := ValidateRepeatParam(repeat); err != nil {
		return "", err
	}

	countedTime, err := CountNextDate(now, date, repeat)
	if err != nil {
		return "", err
	}

	return countedTime.Format(DateFormat), nil
}
