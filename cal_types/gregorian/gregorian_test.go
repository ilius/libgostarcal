package gregorian

import (
	"github.com/ilius/is"
	"testing"

	lib "github.com/ilius/libgostarcal"
)

func TestIsLeap(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[int]bool{
		1990: false,
		1991: false,
		1992: true,
		1993: false,
		1994: false,
		1995: false,
		1996: true,
		1997: false,
		1998: false,
		1999: false,
		2000: true,
		2001: false,
		2002: false,
		2003: false,
		2004: true,
		2005: false,
		2006: false,
		2007: false,
		2008: true,
		2009: false,
		2010: false,
		2011: false,
		2012: true,
		2013: false,
		2014: false,
		2015: false,
		2016: true,
		2017: false,
		2018: false,
		2019: false,
		2020: true,
		2021: false,
		2022: false,
		2023: false,
		2024: true,
		2025: false,
		2026: false,
		2027: false,
		2028: true,
		2029: false,
	}
	for year, isLeap := range testMap {
		is.AddMsg("mismatch isLeap, year=%v", year).Equal(isLeap, IsLeap(year))
	}
}

func TestToJd(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[lib.Date]int{
		{2015, 1, 1}:  2457024,
		{2015, 2, 1}:  2457055,
		{2015, 3, 1}:  2457083,
		{2015, 4, 1}:  2457114,
		{2015, 5, 1}:  2457144,
		{2015, 6, 1}:  2457175,
		{2015, 7, 1}:  2457205,
		{2015, 8, 1}:  2457236,
		{2015, 9, 1}:  2457267,
		{2015, 10, 1}: 2457297,
		{2015, 11, 1}: 2457328,
		{2015, 12, 1}: 2457358,
		{2016, 1, 1}:  2457389,
		{2016, 2, 1}:  2457420,
		{2016, 3, 1}:  2457449,
		{2016, 4, 1}:  2457480,
		{2016, 5, 1}:  2457510,
		{2016, 6, 1}:  2457541,
		{2016, 7, 1}:  2457571,
		{2016, 8, 1}:  2457602,
		{2016, 9, 1}:  2457633,
		{2016, 10, 1}: 2457663,
		{2016, 11, 1}: 2457694,
		{2016, 12, 1}: 2457724,
		{2017, 1, 1}:  2457755,
		{2017, 2, 1}:  2457786,
		{2017, 3, 1}:  2457814,
		{2017, 4, 1}:  2457845,
		{2017, 5, 1}:  2457875,
		{2017, 6, 1}:  2457906,
		{2017, 7, 1}:  2457936,
		{2017, 8, 1}:  2457967,
		{2017, 9, 1}:  2457998,
		{2017, 10, 1}: 2458028,
		{2017, 11, 1}: 2458059,
		{2017, 12, 1}: 2458089,
	}
	for date, jd := range testMap {
		is.AddMsg("mismatch jd, date=%v, jd=%v", date, jd).Equal(jd, ToJd(date))
	}
}

func TestConvert(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	startYear := 1970
	endYear := 2050
	for year := startYear; year < endYear; year++ {
		for month := uint8(1); month <= 12; month++ {
			var monthLen = GetMonthLen(year, month)
			for day := uint8(1); day <= monthLen; day++ {
				var date = lib.Date{year, month, day}
				var jd = ToJd(date)
				var ndate = JdTo(jd)
				is.AddMsg("jd=%v, date=%v, ndate=%v", jd, date, ndate).Equal(date, ndate)
			}
		}
	}
}
