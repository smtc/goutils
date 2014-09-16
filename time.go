package goutils

import (
	"errors"
	"reflect"

	"time"
)

const (
	TIMEFORMAT = "2006-01-02 15:04:05"
)

var (
	TIMEDEFAULT = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	timeType    = reflect.TypeOf(time.Time{})
)

type Time struct {
	time.Time
	f string
}

func (t Time) format() string {
	if t.f == "" {
		t.f = TIMEFORMAT
	}
	return t.Time.Format(t.f)
}

func (t Time) MarshalText() ([]byte, error) {
	return []byte(t.format()), nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJson: year outside of range [0,9999]")
	}
	return []byte(`"` + t.format() + `"`), nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.Parse(`"`+TIMEFORMAT+`"`, string(data))
	return
}

func (t *Time) SetFormat(s string) {
	t.f = s
}
