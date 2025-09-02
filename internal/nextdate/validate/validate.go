package validate_nextdate

import (
	"errors"
	"slices"

	additional_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/additional"
)

func ValidateRepeatParam(repeat string) error {
	firstLetters := []string{"d", "y", "w", "m"}
	splitedRepeat := additional_nextdate.SplitRepeat(repeat)
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
