package validate_nextdate

import (
	"errors"

	additional_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/additional"
)

func ValidateRepeatParam(repeat string) error {
	splitedRepeat := additional_nextdate.SplitRepeat(repeat)
	if len(splitedRepeat) < 1 {
		return errors.New("empty repeat")
	}

	option := splitedRepeat[0]

	if err := validateRepeatbyOption(option, repeat); err != nil {
		return err
	}

	return nil
}
