package nextdate

import (
	"time"

	count_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/count"
	validate_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/validate"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}

	if err := validate_nextdate.ValidateRepeatParam(repeat); err != nil {
		return "", err
	}

	countedTime, err := count_nextdate.CountNextDate(now, date, repeat)
	if err != nil {
		return "", err
	}

	return countedTime.Format(DateFormat), nil
}
