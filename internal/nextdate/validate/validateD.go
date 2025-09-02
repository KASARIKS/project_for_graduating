package validate_nextdate

import (
	"errors"
	"strconv"

	additional_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/additional"
)

func validateRepeatD(repeat string) error {
	splitedRepeat := additional_nextdate.SplitRepeat(repeat)
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
