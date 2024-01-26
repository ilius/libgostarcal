package gregorian

import (
	"testing"

	"github.com/ilius/is/v2"
	"github.com/ilius/libgostarcal/cal_types"

	lib "github.com/ilius/libgostarcal"
)

var calType cal_types.CalType = New()

func TestIsLeap(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[int]string{
		-50: " ",
		-49: " ",
		-48: "L",
		-47: " ",
		-46: " ",
		-45: " ",
		-44: "L",
		-43: " ",
		-42: " ",
		-41: " ",
		-40: "L",
		-39: " ",
		-38: " ",
		-37: " ",
		-36: "L",
		-35: " ",
		-34: " ",
		-33: " ",
		-32: "L",
		-31: " ",
		-30: " ",
		-29: " ",
		-28: "L",
		-27: " ",
		-26: " ",
		-25: " ",
		-24: "L",
		-23: " ",
		-22: " ",
		-21: " ",
		-20: "L",
		-19: " ",
		-18: " ",
		-17: " ",
		-16: "L",
		-15: " ",
		-14: " ",
		-13: " ",
		-12: "L",
		-11: " ",
		-10: " ",
		-9:  " ",
		-8:  "L",
		-7:  " ",
		-6:  " ",
		-5:  " ",
		-4:  "L",
		-3:  " ",
		-2:  " ",
		-1:  " ",
		0:   "L",
		1:   " ",
		2:   " ",
		3:   " ",
		4:   "L",
		5:   " ",
		6:   " ",
		7:   " ",
		8:   "L",
		9:   " ",
		10:  " ",
		11:  " ",
		12:  "L",
		13:  " ",
		14:  " ",
		15:  " ",
		16:  "L",
		17:  " ",
		18:  " ",
		19:  " ",
		20:  "L",
		21:  " ",
		22:  " ",
		23:  " ",
		24:  "L",
		25:  " ",
		26:  " ",
		27:  " ",
		28:  "L",
		29:  " ",
		30:  " ",
		31:  " ",
		32:  "L",
		33:  " ",
		34:  " ",
		35:  " ",
		36:  "L",
		37:  " ",
		38:  " ",
		39:  " ",
		40:  "L",
		41:  " ",
		42:  " ",
		43:  " ",
		44:  "L",
		45:  " ",
		46:  " ",
		47:  " ",
		48:  "L",
		49:  " ",
		50:  " ",

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
		is.AddMsg("mismatch isLeap, year=%v", year).Equal(calType.IsLeap(year), isLeap)
	}
}

func TestToJd(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[*lib.Date]int{
		lib.NewDate(-50, 1, 1): 1702798,
		lib.NewDate(-49, 1, 1): 1703163,
		lib.NewDate(-48, 1, 1): 1703528,
		lib.NewDate(-47, 1, 1): 1703894,
		lib.NewDate(-46, 1, 1): 1704259,
		lib.NewDate(-45, 1, 1): 1704624,
		lib.NewDate(-44, 1, 1): 1704989,
		lib.NewDate(-43, 1, 1): 1705355,
		lib.NewDate(-42, 1, 1): 1705720,
		lib.NewDate(-41, 1, 1): 1706085,
		lib.NewDate(-40, 1, 1): 1706450,
		lib.NewDate(-39, 1, 1): 1706816,
		lib.NewDate(-38, 1, 1): 1707181,
		lib.NewDate(-37, 1, 1): 1707546,
		lib.NewDate(-36, 1, 1): 1707911,
		lib.NewDate(-35, 1, 1): 1708277,
		lib.NewDate(-34, 1, 1): 1708642,
		lib.NewDate(-33, 1, 1): 1709007,
		lib.NewDate(-32, 1, 1): 1709372,
		lib.NewDate(-31, 1, 1): 1709738,
		lib.NewDate(-30, 1, 1): 1710103,
		lib.NewDate(-29, 1, 1): 1710468,
		lib.NewDate(-28, 1, 1): 1710833,
		lib.NewDate(-27, 1, 1): 1711199,
		lib.NewDate(-26, 1, 1): 1711564,
		lib.NewDate(-25, 1, 1): 1711929,
		lib.NewDate(-24, 1, 1): 1712294,
		lib.NewDate(-23, 1, 1): 1712660,
		lib.NewDate(-22, 1, 1): 1713025,
		lib.NewDate(-21, 1, 1): 1713390,
		lib.NewDate(-20, 1, 1): 1713755,
		lib.NewDate(-19, 1, 1): 1714121,
		lib.NewDate(-18, 1, 1): 1714486,
		lib.NewDate(-17, 1, 1): 1714851,
		lib.NewDate(-16, 1, 1): 1715216,
		lib.NewDate(-15, 1, 1): 1715582,
		lib.NewDate(-14, 1, 1): 1715947,
		lib.NewDate(-13, 1, 1): 1716312,
		lib.NewDate(-12, 1, 1): 1716677,
		lib.NewDate(-11, 1, 1): 1717043,
		lib.NewDate(-10, 1, 1): 1717408,
		lib.NewDate(-9, 1, 1):  1717773,
		lib.NewDate(-8, 1, 1):  1718138,
		lib.NewDate(-7, 1, 1):  1718504,
		lib.NewDate(-6, 1, 1):  1718869,
		lib.NewDate(-5, 1, 1):  1719234,
		lib.NewDate(-4, 1, 1):  1719599,
		lib.NewDate(-3, 1, 1):  1719965,
		lib.NewDate(-2, 1, 1):  1720330,
		lib.NewDate(-1, 1, 1):  1720695,
		lib.NewDate(0, 1, 1):   1721060,
		lib.NewDate(1, 1, 1):   1721426,
		lib.NewDate(2, 1, 1):   1721791,
		lib.NewDate(3, 1, 1):   1722156,
		lib.NewDate(4, 1, 1):   1722521,
		lib.NewDate(5, 1, 1):   1722887,
		lib.NewDate(6, 1, 1):   1723252,
		lib.NewDate(7, 1, 1):   1723617,
		lib.NewDate(8, 1, 1):   1723982,
		lib.NewDate(9, 1, 1):   1724348,
		lib.NewDate(10, 1, 1):  1724713,
		lib.NewDate(11, 1, 1):  1725078,
		lib.NewDate(12, 1, 1):  1725443,
		lib.NewDate(13, 1, 1):  1725809,
		lib.NewDate(14, 1, 1):  1726174,
		lib.NewDate(15, 1, 1):  1726539,
		lib.NewDate(16, 1, 1):  1726904,
		lib.NewDate(17, 1, 1):  1727270,
		lib.NewDate(18, 1, 1):  1727635,
		lib.NewDate(19, 1, 1):  1728000,
		lib.NewDate(20, 1, 1):  1728365,
		lib.NewDate(21, 1, 1):  1728731,
		lib.NewDate(22, 1, 1):  1729096,
		lib.NewDate(23, 1, 1):  1729461,
		lib.NewDate(24, 1, 1):  1729826,
		lib.NewDate(25, 1, 1):  1730192,
		lib.NewDate(26, 1, 1):  1730557,
		lib.NewDate(27, 1, 1):  1730922,
		lib.NewDate(28, 1, 1):  1731287,
		lib.NewDate(29, 1, 1):  1731653,
		lib.NewDate(30, 1, 1):  1732018,
		lib.NewDate(31, 1, 1):  1732383,
		lib.NewDate(32, 1, 1):  1732748,
		lib.NewDate(33, 1, 1):  1733114,
		lib.NewDate(34, 1, 1):  1733479,
		lib.NewDate(35, 1, 1):  1733844,
		lib.NewDate(36, 1, 1):  1734209,
		lib.NewDate(37, 1, 1):  1734575,
		lib.NewDate(38, 1, 1):  1734940,
		lib.NewDate(39, 1, 1):  1735305,
		lib.NewDate(40, 1, 1):  1735670,
		lib.NewDate(41, 1, 1):  1736036,
		lib.NewDate(42, 1, 1):  1736401,
		lib.NewDate(43, 1, 1):  1736766,
		lib.NewDate(44, 1, 1):  1737131,
		lib.NewDate(45, 1, 1):  1737497,
		lib.NewDate(46, 1, 1):  1737862,
		lib.NewDate(47, 1, 1):  1738227,
		lib.NewDate(48, 1, 1):  1738592,
		lib.NewDate(49, 1, 1):  1738958,
		lib.NewDate(50, 1, 1):  1739323,

		lib.NewDate(2015, 1, 1):  2457024,
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
		lib.NewDate(2016, 1, 1):  2457389,
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
		lib.NewDate(2017, 1, 1):  2457755,
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

func TestGetMonthLen(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[[2]int]int{
		{2015, 1}:  31,
		{2015, 2}:  28,
		{2015, 3}:  31,
		{2015, 4}:  30,
		{2015, 5}:  31,
		{2015, 6}:  30,
		{2015, 7}:  31,
		{2015, 8}:  31,
		{2015, 9}:  30,
		{2015, 10}: 31,
		{2015, 11}: 30,
		{2015, 12}: 31,
		{2016, 1}:  31,
		{2016, 2}:  29,
		{2016, 3}:  31,
		{2016, 4}:  30,
		{2016, 5}:  31,
		{2016, 6}:  30,
		{2016, 7}:  31,
		{2016, 8}:  31,
		{2016, 9}:  30,
		{2016, 10}: 31,
		{2016, 11}: 30,
		{2016, 12}: 31,
		{2017, 1}:  31,
		{2017, 2}:  28,
		{2017, 3}:  31,
		{2017, 4}:  30,
		{2017, 5}:  31,
		{2017, 6}:  30,
		{2017, 7}:  31,
		{2017, 8}:  31,
		{2017, 9}:  30,
		{2017, 10}: 31,
		{2017, 11}: 30,
		{2017, 12}: 31,
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
