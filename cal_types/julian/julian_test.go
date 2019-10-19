package julian

import (
	"testing"

	"github.com/ilius/is"

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
		is.AddMsg("mismatch isLeap, year=%v", year).Equal(IsLeap(year), isLeap)
	}
}

func TestToJd(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[*lib.Date]int{
		lib.NewDate(-1000, 1, 1): 1355808,
		lib.NewDate(-3, 1, 1):    1719963,
		lib.NewDate(-2, 1, 1):    1720328,
		lib.NewDate(-1, 1, 1):    1720693,

		lib.NewDate(2000, 1, 1): 2451558,
		lib.NewDate(2001, 1, 1): 2451924,
		lib.NewDate(2002, 1, 1): 2452289,
		lib.NewDate(2003, 1, 1): 2452654,
		lib.NewDate(2004, 1, 1): 2453019,
		lib.NewDate(2005, 1, 1): 2453385,
		lib.NewDate(2006, 1, 1): 2453750,
		lib.NewDate(2007, 1, 1): 2454115,
		lib.NewDate(2008, 1, 1): 2454480,
		lib.NewDate(2009, 1, 1): 2454846,
		lib.NewDate(2010, 1, 1): 2455211,
		lib.NewDate(2011, 1, 1): 2455576,
		lib.NewDate(2012, 1, 1): 2455941,
		lib.NewDate(2013, 1, 1): 2456307,
		lib.NewDate(2014, 1, 1): 2456672,
		lib.NewDate(2015, 1, 1): 2457037,
		lib.NewDate(2016, 1, 1): 2457402,
		lib.NewDate(2017, 1, 1): 2457768,
		lib.NewDate(2018, 1, 1): 2458133,
		lib.NewDate(2019, 1, 1): 2458498,
		lib.NewDate(2020, 1, 1): 2458863,
		lib.NewDate(2021, 1, 1): 2459229,
		lib.NewDate(2022, 1, 1): 2459594,
		lib.NewDate(2023, 1, 1): 2459959,
		lib.NewDate(2024, 1, 1): 2460324,
		lib.NewDate(2025, 1, 1): 2460690,
		lib.NewDate(2026, 1, 1): 2461055,
		lib.NewDate(2027, 1, 1): 2461420,
		lib.NewDate(2028, 1, 1): 2461785,
		lib.NewDate(2029, 1, 1): 2462151,
		// lib.NewDate(2015, 1, 1): 2457037,
		lib.NewDate(2015, 2, 1):  2457068,
		lib.NewDate(2015, 3, 1):  2457096,
		lib.NewDate(2015, 4, 1):  2457127,
		lib.NewDate(2015, 5, 1):  2457157,
		lib.NewDate(2015, 6, 1):  2457188,
		lib.NewDate(2015, 7, 1):  2457218,
		lib.NewDate(2015, 8, 1):  2457249,
		lib.NewDate(2015, 9, 1):  2457280,
		lib.NewDate(2015, 10, 1): 2457310,
		lib.NewDate(2015, 11, 1): 2457341,
		lib.NewDate(2015, 12, 1): 2457371,
		// lib.NewDate(2016, 1, 1): 2457402,
		lib.NewDate(2016, 2, 1):  2457433,
		lib.NewDate(2016, 3, 1):  2457462,
		lib.NewDate(2016, 4, 1):  2457493,
		lib.NewDate(2016, 5, 1):  2457523,
		lib.NewDate(2016, 6, 1):  2457554,
		lib.NewDate(2016, 7, 1):  2457584,
		lib.NewDate(2016, 8, 1):  2457615,
		lib.NewDate(2016, 9, 1):  2457646,
		lib.NewDate(2016, 10, 1): 2457676,
		lib.NewDate(2016, 11, 1): 2457707,
		lib.NewDate(2016, 12, 1): 2457737,
		// lib.NewDate(2017, 1, 1): 2457768,
		lib.NewDate(2017, 2, 1):  2457799,
		lib.NewDate(2017, 3, 1):  2457827,
		lib.NewDate(2017, 4, 1):  2457858,
		lib.NewDate(2017, 5, 1):  2457888,
		lib.NewDate(2017, 6, 1):  2457919,
		lib.NewDate(2017, 7, 1):  2457949,
		lib.NewDate(2017, 8, 1):  2457980,
		lib.NewDate(2017, 9, 1):  2458011,
		lib.NewDate(2017, 10, 1): 2458041,
		lib.NewDate(2017, 11, 1): 2458072,
		lib.NewDate(2017, 12, 1): 2458102,
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
				date := lib.NewDate(year, month, day)
				jd := ToJd(date)
				ndate := JdTo(jd)
				is.AddMsg("jd=%v, date=%v, ndate=%v", jd, date, ndate).Equal(ndate, date)
			}
		}
	}
}
