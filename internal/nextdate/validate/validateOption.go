package validate_nextdate

import "errors"

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
