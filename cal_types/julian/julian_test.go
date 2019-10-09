package julian

import (
	"testing"

	"github.com/ilius/is"

	lib "github.com/ilius/libgostarcal"
)

func TestDivmod(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[[2]int][2]int{
		{-20, -10}: {2, 0},
		{20, -10}:  {-2, 0},
		{-20, 10}:  {-2, 0},
		{-23, -10}: {2, -3},
		{-23, 10}:  {-3, 7},
		{-13, 10}:  {-2, 7},
		{12, 10}:   {1, 2},
		{22, 10}:   {2, 2},
	}
	for ab, dm := range testMap {
		a := ab[0]
		b := ab[1]
		div, mod := divmod(a, b)
		is.AddMsg("mismatch div, a=%v, b=%v", a, b).Equal(div, dm[0])
		is.AddMsg("mismatch mod, a=%v, b=%v", a, b).Equal(mod, dm[1])
	}
}

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
	testMap := map[lib.Date]int{
		{-1000, 1, 1}: 1355808,
		{-3, 1, 1}:    1719963,
		{-2, 1, 1}:    1720328,
		{-1, 1, 1}:    1720693,

		{2000, 1, 1}: 2451558,
		{2001, 1, 1}: 2451924,
		{2002, 1, 1}: 2452289,
		{2003, 1, 1}: 2452654,
		{2004, 1, 1}: 2453019,
		{2005, 1, 1}: 2453385,
		{2006, 1, 1}: 2453750,
		{2007, 1, 1}: 2454115,
		{2008, 1, 1}: 2454480,
		{2009, 1, 1}: 2454846,
		{2010, 1, 1}: 2455211,
		{2011, 1, 1}: 2455576,
		{2012, 1, 1}: 2455941,
		{2013, 1, 1}: 2456307,
		{2014, 1, 1}: 2456672,
		{2015, 1, 1}: 2457037,
		{2016, 1, 1}: 2457402,
		{2017, 1, 1}: 2457768,
		{2018, 1, 1}: 2458133,
		{2019, 1, 1}: 2458498,
		{2020, 1, 1}: 2458863,
		{2021, 1, 1}: 2459229,
		{2022, 1, 1}: 2459594,
		{2023, 1, 1}: 2459959,
		{2024, 1, 1}: 2460324,
		{2025, 1, 1}: 2460690,
		{2026, 1, 1}: 2461055,
		{2027, 1, 1}: 2461420,
		{2028, 1, 1}: 2461785,
		{2029, 1, 1}: 2462151,
		// {2015, 1, 1}: 2457037,
		{2015, 2, 1}:  2457068,
		{2015, 3, 1}:  2457096,
		{2015, 4, 1}:  2457127,
		{2015, 5, 1}:  2457157,
		{2015, 6, 1}:  2457188,
		{2015, 7, 1}:  2457218,
		{2015, 8, 1}:  2457249,
		{2015, 9, 1}:  2457280,
		{2015, 10, 1}: 2457310,
		{2015, 11, 1}: 2457341,
		{2015, 12, 1}: 2457371,
		// {2016, 1, 1}: 2457402,
		{2016, 2, 1}:  2457433,
		{2016, 3, 1}:  2457462,
		{2016, 4, 1}:  2457493,
		{2016, 5, 1}:  2457523,
		{2016, 6, 1}:  2457554,
		{2016, 7, 1}:  2457584,
		{2016, 8, 1}:  2457615,
		{2016, 9, 1}:  2457646,
		{2016, 10, 1}: 2457676,
		{2016, 11, 1}: 2457707,
		{2016, 12, 1}: 2457737,
		// {2017, 1, 1}: 2457768,
		{2017, 2, 1}:  2457799,
		{2017, 3, 1}:  2457827,
		{2017, 4, 1}:  2457858,
		{2017, 5, 1}:  2457888,
		{2017, 6, 1}:  2457919,
		{2017, 7, 1}:  2457949,
		{2017, 8, 1}:  2457980,
		{2017, 9, 1}:  2458011,
		{2017, 10, 1}: 2458041,
		{2017, 11, 1}: 2458072,
		{2017, 12, 1}: 2458102,
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
