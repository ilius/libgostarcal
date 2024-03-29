package utils

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	. "github.com/ilius/libgostarcal/utils/mapset"

	lib "github.com/ilius/libgostarcal"
	"github.com/ilius/libgostarcal/cal_types"
	"github.com/ilius/libgostarcal/cal_types/gregorian"
)

var greg = gregorian.New()

func init() {
	fmt.Printf("")
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Int64ListToIntList(list []int64) []int {
	list2 := make([]int, len(list))
	for i, v := range list {
		list2[i] = int(v)
	}
	return list2
}

func ParseIntList(str string) ([]int, error) {
	parts := strings.Split(str, " ")
	intParts := make([]int, len(parts))
	for index, part := range parts {
		intPart, err := strconv.ParseInt(part, 10, 0)
		if err != nil {
			return []int{}, err
		}
		intParts[index] = int(intPart)
	}
	return intParts, nil
}

func MonthListIsValid(list []int) bool {
	for _, v := range list {
		if !(v > 0 && v < 13) {
			return false
		}
	}
	return true
}

func DayListIsValid(list []int) bool {
	for _, v := range list {
		if !(v > 0 && v < 40) {
			return false
		}
	}
	return true
}

func WeekDayListIsValid(list []int) bool {
	for _, v := range list {
		if !(v >= 0 && v < 7) {
			return false
		}
	}
	return true
}

func IntListBySet(set Set) []int {
	ret := make([]int, set.Cardinality())
	for index, inf := range set.ToSlice() {
		ret[index] = inf.(int)
	}
	return ret
}

func bisectLeftRange(a []int, v int, lo, hi int) int {
	s := a[lo:hi]
	return sort.Search(len(s), func(i int) bool {
		return s[i] >= v
	})
}

func BisectLeft(a []int, v int) int {
	return bisectLeftRange(a, v, 0, len(a))
}

// tested
func GetUtcOffsetByGDate(gdate lib.Date, loc *time.Location) int {
	t := time.Date(
		gdate.Year,
		time.Month(gdate.Month), // gdate.Month is int
		int(gdate.Day),
		0,   // hour
		0,   // min
		0,   // sec
		0,   // nsec
		loc, // location
	)
	_, offset := t.Zone() // zoneName, offset
	return offset
}

// tested
func GetUtcOffsetByEpoch(epoch int64, loc *time.Location) int {
	// is this working perfectly? FIXME
	// python code is too tricky
	t := time.Unix(epoch, 0).In(loc) // .In useful? FIXME
	_, offset := t.Zone()            // zoneName, offset
	return offset
}

func GetUtcOffsetCurrent(loc *time.Location) int {
	t := time.Now().In(loc)
	_, offset := t.Zone() // zoneName, offset
	return offset
}

// tested
func GetEpochByGDate(gdate *lib.Date, loc *time.Location) int64 {
	t := time.Date(
		gdate.Year,
		time.Month(int(gdate.Month)), // gdate.Month is uint8
		int(gdate.Day),
		0,   // hour
		0,   // min
		0,   // sec
		0,   // nsec
		loc, // location
	)
	return t.Unix()
}

// tested
func GetEpochByJd(jd int, loc *time.Location) int64 {
	return GetEpochByGDate(greg.JdTo(jd), loc)
}

/*
func GetEpochByJd2(jd int, loc *time.Location) int64 {
    localEpoch := int64((jd-J1970) * 86400)
    offset := GetUtcOffsetByGDate(gdate, loc)
    epoch := localEpoch - offset
    offset2 := GetUtcOffsetByEpoch(epoch, loc)
    if offset2 != offset {
        fmt.Println("Warning: GetEpochByJd: offset mistmatch: delta =", offset2-offset, ", gdate =", gdate)
        epoch = localEpoch - offset2
        //3600 seconds error in days when DST is just changed
        //gdate = {2016 9 21}
        //gdate = {2017 9 22}
        //gdate = {2018 9 22}
        //gdate = {2019 9 22}
    }
    return epoch
}*/

func GetFloatJdByEpoch(epoch int64, loc *time.Location) float64 {
	offset := GetUtcOffsetByEpoch(epoch, loc)
	return float64(J1970) + float64(epoch+int64(offset))/86400.0
}

func GetJdByEpoch(epoch int64, loc *time.Location) int {
	return int(math.Floor(GetFloatJdByEpoch(epoch, loc)))
}

// RoundEpochToDay // not useful

func GetJdRangeFromEpochRange(startEpoch int64, endEpoch int64, loc *time.Location) (int, int) {
	startJd := GetJdByEpoch(startEpoch, loc)
	endJd := GetJdByEpoch(endEpoch-1, loc) + 1
	return startJd, endJd
}

func GetHmsBySeconds(second uint) lib.HMS {
	return lib.HMS{
		Hour:   uint8(second / 3600),
		Minute: uint8(second / 60),
		Second: uint8(second % 60), // safe %
	}
}

func GetJhmsByEpoch(epoch int64, loc *time.Location) (int, lib.HMS) {
	// return (jd, hour, minute, second)
	t := time.Unix(epoch, 0).In(loc) // .In useful? FIXME
	jd := greg.ToJd(lib.NewDate(
		t.Year(),
		uint8(t.Month()),
		uint8(t.Day()),
	))
	return jd, lib.HMS{
		Hour:   uint8(t.Hour()),
		Minute: uint8(t.Minute()),
		Second: uint8(t.Second()),
	}
}

func GetEpochByJhms(jd int, hms lib.HMS, loc *time.Location) int64 {
	gdate := greg.JdTo(jd)
	t := time.Date(
		gdate.Year,
		time.Month(gdate.Month), // gdate.Month is uint8
		int(gdate.Day),
		int(hms.Hour),
		int(hms.Minute),
		int(hms.Second),
		0,   // nsec
		loc, // location
	)
	return t.Unix()
}

func GetJdAndSecondsFromEpoch(epoch int64, loc *time.Location) (int, int) {
	// return a tuple (julain_day, extra_seconds) from epoch
	jd, hms := GetJhmsByEpoch(epoch, loc)
	return jd, hms.GetTotalSeconds()
}

func GetCurrentDate(calTypeName string) (*lib.Date, error) {
	t := time.Now() // .In(loc)
	if calTypeName == "gregorian" {
		return lib.NewDate(t.Year(), uint8(t.Month()), uint8(t.Day())), nil
	}
	calType, ok := cal_types.CalTypesMap[calTypeName]
	if !ok {
		return nil,
			errors.New("invalid calendar type '" + calTypeName + "'")
	}
	loc := t.Location() // FIXME
	jd := GetJdByEpoch(t.Unix(), loc)
	return calType.JdTo(jd), nil
}
