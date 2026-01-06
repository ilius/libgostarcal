package ethiopian

import (
	"testing"

	"github.com/ilius/is/v2"
	"github.com/ilius/libgostarcal/cal_types"

	lib "github.com/ilius/libgostarcal"
)

var calType cal_types.CalType = New()

func TestIsLeap(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[int]bool{
		1990: false,
		1991: true,
		1992: false,
		1993: false,
		1994: false,
		1995: true,
		1996: false,
		1997: false,
		1998: false,
		1999: true,
		2000: false,
		2001: false,
		2002: false,
		2003: true,
		2004: false,
		2005: false,
		2006: false,
		2007: true,
		2008: false,
		2009: false,
		2010: false,
		2011: true,
		2012: false,
		2013: false,
		2014: false,
		2015: true,
		2016: false,
		2017: false,
		2018: false,
		2019: true,
		2020: false,
		2021: false,
		2022: false,
		2023: true,
		2024: false,
		2025: false,
		2026: false,
		2027: true,
		2028: false,
		2029: false,
	}
	for year, isLeap := range testMap {
		is.AddMsg("mismatch isLeap, year=%v", year).Equal(calType.IsLeap(year), isLeap)
	}
}

func TestToJd(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[*lib.Date]int{
		lib.NewDate(2015, 1, 1):  2459834,
		lib.NewDate(2015, 2, 1):  2459864,
		lib.NewDate(2015, 3, 1):  2459894,
		lib.NewDate(2015, 4, 1):  2459924,
		lib.NewDate(2015, 5, 1):  2459954,
		lib.NewDate(2015, 6, 1):  2459984,
		lib.NewDate(2015, 7, 1):  2460014,
		lib.NewDate(2015, 8, 1):  2460044,
		lib.NewDate(2015, 9, 1):  2460074,
		lib.NewDate(2015, 10, 1): 2460104,
		lib.NewDate(2015, 11, 1): 2460134,
		lib.NewDate(2015, 12, 1): 2460164,
		lib.NewDate(2016, 1, 1):  2460200,
		lib.NewDate(2016, 2, 1):  2460230,
		lib.NewDate(2016, 3, 1):  2460260,
		lib.NewDate(2016, 4, 1):  2460290,
		lib.NewDate(2016, 5, 1):  2460320,
		lib.NewDate(2016, 6, 1):  2460350,
		lib.NewDate(2016, 7, 1):  2460380,
		lib.NewDate(2016, 8, 1):  2460410,
		lib.NewDate(2016, 9, 1):  2460440,
		lib.NewDate(2016, 10, 1): 2460470,
		lib.NewDate(2016, 11, 1): 2460500,
		lib.NewDate(2016, 12, 1): 2460530,
		lib.NewDate(2017, 1, 1):  2460565,
		lib.NewDate(2017, 2, 1):  2460595,
		lib.NewDate(2017, 3, 1):  2460625,
		lib.NewDate(2017, 4, 1):  2460655,
		lib.NewDate(2017, 5, 1):  2460685,
		lib.NewDate(2017, 6, 1):  2460715,
		lib.NewDate(2017, 7, 1):  2460745,
		lib.NewDate(2017, 8, 1):  2460775,
		lib.NewDate(2017, 9, 1):  2460805,
		lib.NewDate(2017, 10, 1): 2460835,
		lib.NewDate(2017, 11, 1): 2460865,
		lib.NewDate(2017, 12, 1): 2460895,
	}
	for date, jd := range testMap {
		is.AddMsg("mismatch jd, date=%v, jd=%v", date, jd).Equal(calType.ToJd(date), jd)
	}
}

func TestGetMonthLen(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[[2]int]uint8{
		{2015, 1}:  30,
		{2015, 2}:  30,
		{2015, 3}:  30,
		{2015, 4}:  30,
		{2015, 5}:  30,
		{2015, 6}:  30,
		{2015, 7}:  30,
		{2015, 8}:  30,
		{2015, 9}:  30,
		{2015, 10}: 30,
		{2015, 11}: 30,
		{2015, 12}: 36,
		{2016, 1}:  30,
		{2016, 2}:  30,
		{2016, 3}:  30,
		{2016, 4}:  30,
		{2016, 5}:  30,
		{2016, 6}:  30,
		{2016, 7}:  30,
		{2016, 8}:  30,
		{2016, 9}:  30,
		{2016, 10}: 30,
		{2016, 11}: 30,
		{2016, 12}: 35,
		{2017, 1}:  30,
		{2017, 2}:  30,
		{2017, 3}:  30,
		{2017, 4}:  30,
		{2017, 5}:  30,
		{2017, 6}:  30,
		{2017, 7}:  30,
		{2017, 8}:  30,
		{2017, 9}:  30,
		{2017, 10}: 30,
		{2017, 11}: 30,
		{2017, 12}: 35,
	}
	for ym, monthLen := range testMap {
		year := ym[0]
		month := uint8(ym[1])
		is.AddMsg("ym={%v, %v}", year, month).Equal(calType.GetMonthLen(year, month), monthLen)
	}
}

func TestConvert(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	startYear := 1970
	endYear := 2050
	for year := startYear; year < endYear; year++ {
		for month := uint8(1); month <= 12; month++ {
			monthLen := calType.GetMonthLen(year, month)
			for day := uint8(1); day <= monthLen; day++ {
				date := lib.NewDate(year, month, day)
				jd := calType.ToJd(date)
				ndate := calType.JdTo(jd)
				is.AddMsg("jd=%v, date=%v, ndate=%v", jd, date, ndate).Equal(ndate, date)
			}
		}
	}
}
