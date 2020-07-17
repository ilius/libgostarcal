package libscal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func NewDate(year int, month uint8, day uint8) *Date {
	return &Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

type Date struct {
	Year  int
	Month uint8
	Day   uint8
}

func (date *Date) StringWithSep(sep string) string {
	return fmt.Sprintf(
		"%.4d%s%.2d%s%.2d",
		date.Year,
		sep,
		date.Month,
		sep,
		date.Day,
	)
}

func (date *Date) String() string {
	return fmt.Sprintf("%.4d/%.2d/%.2d", date.Year, date.Month, date.Day)
}

func (date *Date) Repr() string {
	return fmt.Sprintf("NewDate(%d, %d, %d)", date.Year, date.Month, date.Day)
}

func (date *Date) IsValid() bool {
	return date.Month > 0 && date.Month < 13 && date.Day > 0 && date.Day < 40
}

func ParseDate(str string) (*Date, error) {
	parts := strings.Split(str, "/")
	if len(parts) != 3 {
		return nil,
			errors.New("invalid Date string '" + str + "'")
	}
	var err error
	var y, m, d int64
	y, err = strconv.ParseInt(parts[0], 10, 0)
	if err != nil {
		return nil, err
	}
	m, err = strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		return nil, err
	}
	d, err = strconv.ParseInt(parts[2], 10, 0)
	if err != nil {
		return nil, err
	}
	return NewDate(int(y), uint8(m), uint8(d)), nil
}

func ParseDateList(str string) ([]*Date, error) {
	parts := strings.Split(str, " ")
	dates := make([]*Date, len(parts))
	for index, part := range parts {
		date, err := ParseDate(part)
		if err != nil {
			return nil, err
		}
		dates[index] = date
	}
	return dates, nil
}

type DateHMS struct {
	*Date
	*HMS
}

func (dt DateHMS) String() string {
	return dt.Date.String() + " " + dt.HMS.String()
}

func (dt DateHMS) Repr() string {
	return fmt.Sprintf("DateHMS{Date: %s, HMS: %s}", dt.Date.Repr(), dt.HMS.Repr())
}

func (dt DateHMS) IsValid() bool {
	return dt.Date.IsValid() && dt.HMS.IsValid()
}

func ParseDateHMS(str string) (*DateHMS, error) {
	parts := strings.Split(str, " ")
	if len(parts) != 2 {
		return nil,
			errors.New("invalid DateHMS string '" + str + "'")
	}
	date, err := ParseDate(parts[0])
	if err != nil {
		return nil, err
	}
	hms, err := ParseHMS(parts[1])
	if err != nil {
		return nil, err
	}
	return &DateHMS{
		Date: date,
		HMS:  hms,
	}, nil
}
