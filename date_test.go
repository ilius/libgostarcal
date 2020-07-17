package libscal

import (
	"testing"

	"github.com/ilius/is/v2"
)

func TestDateStringWithSep(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	is.Equal(
		NewDate(2019, 12, 1).StringWithSep("."),
		"2019.12.01",
	)
	is.Equal(
		NewDate(2021, 1, 3).StringWithSep(","),
		"2021,01,03",
	)
}

func TestDateString(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	is.Equal(
		NewDate(2021, 1, 3).String(),
		"2021/01/03",
	)
}

func TestDateRepr(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	is.Equal(
		NewDate(2021, 1, 3).Repr(),
		"NewDate(2021, 1, 3)",
	)
}

func TestDateIsValid(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	test := func(valid bool, year int, month uint8, day uint8) {
		date := NewDate(year, month, day)
		is.Msg("%v", date).Equal(date.IsValid(), valid)
	}
	test(false, 2020, 0, 0)
	test(false, 2020, 0, 1)
	test(false, 2020, 1, 0)
	test(false, 2020, 13, 1)
	test(false, 2020, 1, 40)
	test(true, 2017, 10, 25)
	test(true, 2013, 10, 20)
	test(true, 2016, 8, 23)
	test(true, 2010, 1, 5)
	test(true, 2010, 10, 19)
	test(true, 2010, 7, 3)
	test(true, 2017, 3, 23)
	test(true, 2015, 6, 16)
	test(true, 2019, 6, 23)
	test(true, 2019, 10, 17)
	test(true, 2016, 9, 14)
	test(true, 2010, 12, 15)
	test(true, 2014, 7, 29)
	test(true, 2011, 11, 26)
	test(true, 2019, 2, 9)
	test(true, 2011, 7, 29)
	test(true, 2015, 5, 25)
	test(true, 2011, 12, 16)
	test(true, 2018, 1, 18)
	test(true, 2010, 8, 9)
}

func TestParseDate(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	test := func(dateStr string, correctDateStr string, errMsg string) {
		is := is.Msg("%v", dateStr)
		date, err := ParseDate(dateStr)
		if errMsg != "" {
			is.ErrMsg(err, errMsg)
			return
		}
		is.NotErr(err)
		is.Equal(date.String(), correctDateStr)
	}
	test("2016/3/1", "2016/03/01", "")
	test("2016/12/1", "2016/12/01", "")
	test("2016/12/01", "2016/12/01", "")
	test("2016/04/27", "2016/04/27", "")
	test("2015-12-21", "", "invalid Date string '2015-12-21'")
}

func TestParseDateList(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	test := func(dateListStr string, errMsg string, expectedStrList ...string) {
		dateList, err := ParseDateList(dateListStr)
		if errMsg != "" {
			is.Nil(dateList)
			is.ErrMsg(err, errMsg)
			return
		}
		is.NotErr(err)
		is.Equal(len(dateList), len(expectedStrList))
		for i := 0; i < len(dateList); i++ {
			is.Equal(dateList[i].String(), expectedStrList[i])
		}
	}
	test(
		"2014/08/05,2010/07/11",
		"invalid Date string '2014/08/05,2010/07/11'",
	)
	test(
		"2014/08/05",
		"", // no error
		"2014/08/05",
	)
	test(
		"2014/08/05 2010/07/11 2019/06/19 2016/08/08",
		"", // no error
		"2014/08/05",
		"2010/07/11",
		"2019/06/19",
		"2016/08/08",
	)
}

func TestDateHMSString(t *testing.T) {
	is := is.New(t)
	x := &DateHMS{
		Date: NewDate(2010, 01, 03),
		HMS:  NewHMS(7, 8, 9),
	}
	is.Equal(x.String(), "2010/01/03 07:08:09")
}

func TestDateHMSRepr(t *testing.T) {
	is := is.New(t)
	x := &DateHMS{
		Date: NewDate(2010, 01, 03),
		HMS:  NewHMS(7, 8, 9),
	}
	is.Equal(x.Repr(), "DateHMS{Date: NewDate(2010, 1, 3), HMS: NewHMS(7, 8, 9)}")
}

func TestDateHMSIsValid(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	test := func(
		valid bool,
		year int, month uint8, day uint8,
		hour uint8, minute uint8, second uint8,
	) {
		x := &DateHMS{
			Date: NewDate(year, month, day),
			HMS:  NewHMS(hour, minute, second),
		}
		is.Msg("%v", x).Equal(x.IsValid(), valid)
	}
	test(false, 2020, 0, 0, 0, 0, 0)
	test(false, 2020, 0, 1, 0, 0, 0)
	test(false, 2020, 1, 0, 0, 0, 0)
	test(false, 2020, 13, 1, 0, 0, 0)
	test(false, 2020, 1, 40, 0, 0, 0)
	test(false, 2020, 1, 1, 24, 0, 0)
	test(false, 2020, 1, 1, 0, 60, 0)
	test(false, 2020, 1, 1, 0, 0, 60)
	test(true, 2017, 10, 25, 23, 59, 59)
	test(true, 2013, 10, 20, 10, 12, 13)
	test(true, 2016, 8, 23, 14, 1, 0)
	test(true, 2010, 1, 5, 15, 55, 0)
	test(true, 2010, 10, 19, 0, 0, 1)
}

func TestParseDateHMS(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	test := func(str string, correctStr string, errMsg string) {
		is := is.Msg("%v", str)
		x, err := ParseDateHMS(str)
		if errMsg != "" {
			is.ErrMsg(err, errMsg)
			return
		}
		is.NotErr(err)
		is.Equal(x.String(), correctStr)
	}
	test("2016/3/1 12:30:0", "2016/03/01 12:30:00", "")
	test("2016/12/1 0:0:0", "2016/12/01 00:00:00", "")
	test("2016/12/01 1:2:3", "2016/12/01 01:02:03", "")
	test("2016/04/27", "", "invalid DateHMS string '2016/04/27'")
	test("2015-12-21 0:0:0", "", "invalid Date string '2015-12-21'")
}
