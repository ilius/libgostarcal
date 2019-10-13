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
		{-50, 1, 1}: 1702798,
		{-49, 1, 1}: 1703163,
		{-48, 1, 1}: 1703528,
		{-47, 1, 1}: 1703894,
		{-46, 1, 1}: 1704259,
		{-45, 1, 1}: 1704624,
		{-44, 1, 1}: 1704989,
		{-43, 1, 1}: 1705355,
		{-42, 1, 1}: 1705720,
		{-41, 1, 1}: 1706085,
		{-40, 1, 1}: 1706450,
		{-39, 1, 1}: 1706816,
		{-38, 1, 1}: 1707181,
		{-37, 1, 1}: 1707546,
		{-36, 1, 1}: 1707911,
		{-35, 1, 1}: 1708277,
		{-34, 1, 1}: 1708642,
		{-33, 1, 1}: 1709007,
		{-32, 1, 1}: 1709372,
		{-31, 1, 1}: 1709738,
		{-30, 1, 1}: 1710103,
		{-29, 1, 1}: 1710468,
		{-28, 1, 1}: 1710833,
		{-27, 1, 1}: 1711199,
		{-26, 1, 1}: 1711564,
		{-25, 1, 1}: 1711929,
		{-24, 1, 1}: 1712294,
		{-23, 1, 1}: 1712660,
		{-22, 1, 1}: 1713025,
		{-21, 1, 1}: 1713390,
		{-20, 1, 1}: 1713755,
		{-19, 1, 1}: 1714121,
		{-18, 1, 1}: 1714486,
		{-17, 1, 1}: 1714851,
		{-16, 1, 1}: 1715216,
		{-15, 1, 1}: 1715582,
		{-14, 1, 1}: 1715947,
		{-13, 1, 1}: 1716312,
		{-12, 1, 1}: 1716677,
		{-11, 1, 1}: 1717043,
		{-10, 1, 1}: 1717408,
		{-9, 1, 1}:  1717773,
		{-8, 1, 1}:  1718138,
		{-7, 1, 1}:  1718504,
		{-6, 1, 1}:  1718869,
		{-5, 1, 1}:  1719234,
		{-4, 1, 1}:  1719599,
		{-3, 1, 1}:  1719965,
		{-2, 1, 1}:  1720330,
		{-1, 1, 1}:  1720695,
		{0, 1, 1}:   1721060,
		{1, 1, 1}:   1721426,
		{2, 1, 1}:   1721791,
		{3, 1, 1}:   1722156,
		{4, 1, 1}:   1722521,
		{5, 1, 1}:   1722887,
		{6, 1, 1}:   1723252,
		{7, 1, 1}:   1723617,
		{8, 1, 1}:   1723982,
		{9, 1, 1}:   1724348,
		{10, 1, 1}:  1724713,
		{11, 1, 1}:  1725078,
		{12, 1, 1}:  1725443,
		{13, 1, 1}:  1725809,
		{14, 1, 1}:  1726174,
		{15, 1, 1}:  1726539,
		{16, 1, 1}:  1726904,
		{17, 1, 1}:  1727270,
		{18, 1, 1}:  1727635,
		{19, 1, 1}:  1728000,
		{20, 1, 1}:  1728365,
		{21, 1, 1}:  1728731,
		{22, 1, 1}:  1729096,
		{23, 1, 1}:  1729461,
		{24, 1, 1}:  1729826,
		{25, 1, 1}:  1730192,
		{26, 1, 1}:  1730557,
		{27, 1, 1}:  1730922,
		{28, 1, 1}:  1731287,
		{29, 1, 1}:  1731653,
		{30, 1, 1}:  1732018,
		{31, 1, 1}:  1732383,
		{32, 1, 1}:  1732748,
		{33, 1, 1}:  1733114,
		{34, 1, 1}:  1733479,
		{35, 1, 1}:  1733844,
		{36, 1, 1}:  1734209,
		{37, 1, 1}:  1734575,
		{38, 1, 1}:  1734940,
		{39, 1, 1}:  1735305,
		{40, 1, 1}:  1735670,
		{41, 1, 1}:  1736036,
		{42, 1, 1}:  1736401,
		{43, 1, 1}:  1736766,
		{44, 1, 1}:  1737131,
		{45, 1, 1}:  1737497,
		{46, 1, 1}:  1737862,
		{47, 1, 1}:  1738227,
		{48, 1, 1}:  1738592,
		{49, 1, 1}:  1738958,
		{50, 1, 1}:  1739323,

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
