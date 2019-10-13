package gregorian

import (
	"testing"

	"github.com/ilius/is"

	lib "github.com/ilius/libgostarcal"
)

func TestIsLeap(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[int]string{
		1990: " ",
		1991: " ",
		1992: "L",
		1993: " ",
		1994: " ",
		1995: " ",
		1996: "L",
		1997: " ",
		1998: " ",
		1999: " ",
		2000: "L",
		2001: " ",
		2002: " ",
		2003: " ",
		2004: "L",
		2005: " ",
		2006: " ",
		2007: " ",
		2008: "L",
		2009: " ",
		2010: " ",
		2011: " ",
		2012: "L",
		2013: " ",
		2014: " ",
		2015: " ",
		2016: "L",
		2017: " ",
		2018: " ",
		2019: " ",
		2020: "L",
		2021: " ",
		2022: " ",
		2023: " ",
		2024: "L",
		2025: " ",
		2026: " ",
		2027: " ",
		2028: "L",
		2029: " ",
	}
	for year, isLeapStr := range testMap {
		isLeap := isLeapStr == "L"
		is.AddMsg("mismatch isLeap, year=%v", year).Equal(IsLeap(year), isLeap)
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
		is.AddMsg("mismatch jd, date=%v, jd=%v", date, jd).Equal(ToJd(date), jd)
	}
}

func TestConvert(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	startYear := 1970
	endYear := 2050
	for year := startYear; year < endYear; year++ {
		for month := uint8(1); month <= 12; month++ {
			monthLen := GetMonthLen(year, month)
			for day := uint8(1); day <= monthLen; day++ {
				date := lib.Date{year, month, day}
				jd := ToJd(date)
				ndate := JdTo(jd)
				is.AddMsg("jd=%v, date=%v, ndate=%v", jd, date, ndate).Equal(ndate, date)
			}
		}
	}
}
