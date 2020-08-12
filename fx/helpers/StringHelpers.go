package helpers

import (
	"strconv"
	"time"
)

const (
	PAD_TYPE_LEFT  = 0
	PAD_TYPE_RIGHT = 1
)

func StringPad(needle string, padString string, nPad int, padType int) string {

	if padType == PAD_TYPE_LEFT {
		str := ""
		nPad = nPad - len(needle)
		for i := 0; i < nPad; i++ {
			str += padString
		}
		return str + needle

	} else if padType == PAD_TYPE_RIGHT {

		str := ""
		nPad = nPad - len(needle)
		for i := 0; i < nPad; i++ {
			str += padString
		}
		return needle + str

	} else {
		panic("invalid padType=" + strconv.Itoa(padType))
	}
}

func StrToDate(strDate string) (*time.Time, error) {

	if strDate == "" {
		return nil, nil

	} else {

		layout := "2006-01-02"
		tDate, err := time.Parse(layout, strDate)

		if err != nil {
			return nil, err
		}

		return &tDate, nil

	}
}

func StrToTime(strTime string) (*time.Time, error) {

	if strTime == "" {
		return nil, nil

	} else {

		layout := "15:04"
		tDate, err := time.Parse(layout, strTime)
		if err != nil {
			return nil, err
		}
		return &tDate, nil

	}
}
