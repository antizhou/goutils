package time

import (
	"time"
)

func UnixTsFormat(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

func UnixTsOfTodayStart(ts int64) int64 {
	tms := time.Unix(ts, 0)
	hour := tms.Hour()
	min := tms.Minute()
	sec := tms.Second()
	return tms.Unix() - int64(hour*60*60) - int64(min*60) - int64(sec)
}
