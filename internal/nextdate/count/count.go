package count_nextdate

import (
	"time"

	additional_nextdate "github.com/kasariks/project_for_graduating/internal/nextdate/additional"
)

func CountNextDate(now time.Time, dstartTime time.Time, repeat string) (time.Time, error) {
	splitedRepeat := additional_nextdate.SplitRepeat(repeat)
	option := splitedRepeat[0]
	var err error

	switch option {
	case "d":
		dstartTime, err = CountForDOption(now, dstartTime, splitedRepeat)
	case "y":
		dstartTime, err = CountForYOption(now, dstartTime)
	case "w":
		dstartTime, err = CountForWOption(now, dstartTime, splitedRepeat)
	case "m":
		dstartTime, err = CountForMOption(now, dstartTime, repeat)
	}

	return dstartTime, err
}
