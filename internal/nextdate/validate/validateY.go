package validate_nextdate

import (
	"errors"

	additional_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/additional"
)

func validateRepeatY(repeat string) error {
	splitedRepeat := additional_nextdate.SplitRepeat(repeat)
	if len(splitedRepeat) > 1 {
		return errors.New("unsupported format length for y option")
	}

	return nil
}
