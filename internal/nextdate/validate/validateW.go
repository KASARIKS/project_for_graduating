package validate_nextdate

import (
	"errors"
	"strconv"

	additional_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/additional"
)

func validateRepeatW(repeat string) error {
	splitedRepeat := additional_nextdate.SplitRepeat(repeat)
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
