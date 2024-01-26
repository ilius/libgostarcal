package gregorian_proleptic

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
		is.AddMsg("mismatch isLeap, year=%v", year).Equal(calType.IsLeap(year), isLeap)
	}
}

func TestToJd(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[*lib.Date]int{
		lib.NewDate(2000, 1, 1):  2451545,
		lib.NewDate(2001, 1, 1):  2451911,
		lib.NewDate(2002, 1, 1):  2452276,
		lib.NewDate(2003, 1, 1):  2452641,
		lib.NewDate(2004, 1, 1):  2453006,
		lib.NewDate(2005, 1, 1):  2453372,
		lib.NewDate(2006, 1, 1):  2453737,
		lib.NewDate(2007, 1, 1):  2454102,
		lib.NewDate(2008, 1, 1):  2454467,
		lib.NewDate(2009, 1, 1):  2454833,
		lib.NewDate(2010, 1, 1):  2455198,
		lib.NewDate(2011, 1, 1):  2455563,
		lib.NewDate(2012, 1, 1):  2455928,
		lib.NewDate(2013, 1, 1):  2456294,
		lib.NewDate(2014, 1, 1):  2456659,
		lib.NewDate(2015, 1, 1):  2457024,
		lib.NewDate(2016, 1, 1):  2457389,
		lib.NewDate(2017, 1, 1):  2457755,
		lib.NewDate(2018, 1, 1):  2458120,
		lib.NewDate(2019, 1, 1):  2458485,
		lib.NewDate(2020, 1, 1):  2458850,
		lib.NewDate(2021, 1, 1):  2459216,
		lib.NewDate(2022, 1, 1):  2459581,
		lib.NewDate(2023, 1, 1):  2459946,
		lib.NewDate(2024, 1, 1):  2460311,
		lib.NewDate(2025, 1, 1):  2460677,
		lib.NewDate(2026, 1, 1):  2461042,
		lib.NewDate(2027, 1, 1):  2461407,
		lib.NewDate(2028, 1, 1):  2461772,
		lib.NewDate(2029, 1, 1):  2462138,
		lib.NewDate(2015, 2, 1):  2457055,
		lib.NewDate(2015, 3, 1):  2457083,
		lib.NewDate(2015, 4, 1):  2457114,
		lib.NewDate(2015, 5, 1):  2457144,
		lib.NewDate(2015, 6, 1):  2457175,
		lib.NewDate(2015, 7, 1):  2457205,
		lib.NewDate(2015, 8, 1):  2457236,
		lib.NewDate(2015, 9, 1):  2457267,
		lib.NewDate(2015, 10, 1): 2457297,
		lib.NewDate(2015, 11, 1): 2457328,
		lib.NewDate(2015, 12, 1): 2457358,
		lib.NewDate(2016, 2, 1):  2457420,
		lib.NewDate(2016, 3, 1):  2457449,
		lib.NewDate(2016, 4, 1):  2457480,
		lib.NewDate(2016, 5, 1):  2457510,
		lib.NewDate(2016, 6, 1):  2457541,
		lib.NewDate(2016, 7, 1):  2457571,
		lib.NewDate(2016, 8, 1):  2457602,
		lib.NewDate(2016, 9, 1):  2457633,
		lib.NewDate(2016, 10, 1): 2457663,
		lib.NewDate(2016, 11, 1): 2457694,
		lib.NewDate(2016, 12, 1): 2457724,
		lib.NewDate(2017, 2, 1):  2457786,
		lib.NewDate(2017, 3, 1):  2457814,
		lib.NewDate(2017, 4, 1):  2457845,
		lib.NewDate(2017, 5, 1):  2457875,
		lib.NewDate(2017, 6, 1):  2457906,
		lib.NewDate(2017, 7, 1):  2457936,
		lib.NewDate(2017, 8, 1):  2457967,
		lib.NewDate(2017, 9, 1):  2457998,
		lib.NewDate(2017, 10, 1): 2458028,
		lib.NewDate(2017, 11, 1): 2458059,
		lib.NewDate(2017, 12, 1): 2458089,
	}
	for date, jd := range testMap {
		is.AddMsg("mismatch jd, date=%v, jd=%v", date, jd).Equal(calType.ToJd(date), jd)
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
