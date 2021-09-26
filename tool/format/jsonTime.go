package format

import (
	"fmt"
	"time"
)

type JSONTime struct {
	time.Time
}

const (
	TimeFormat = "2006-01-02 15:04:05"
)

// 格式化时区时间
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(formatted), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = JSONTime{now}
	return
}
