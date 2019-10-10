package jalali

import (
	"testing"

	"github.com/ilius/is"

	lib "github.com/ilius/libgostarcal"
)

func TestIsLeap(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[int]string{
		1360: " ",
		1361: " ",
		1362: "L",
		1363: " ",
		1364: " ",
		1365: " ",
		1366: "L",
		1367: " ",
		1368: " ",
		1369: " ",
		1370: "L",
		1371: " ",
		1372: " ",
		1373: " ",
		1374: " ",
		1375: "L",
		1376: " ",
		1377: " ",
		1378: " ",
		1379: "L",
		1380: " ",
		1381: " ",
		1382: " ",
		1383: "L",
		1384: " ",
		1385: " ",
		1386: " ",
		1387: "L",
		1388: " ",
		1389: " ",
		1390: " ",
		1391: "L",
		1392: " ",
		1393: " ",
		1394: " ",
		1395: "L",
		1396: " ",
		1397: " ",
		1398: " ",
		1399: "L",
		1400: " ",
		1401: " ",
		1402: " ",
		1404: "L",
		1405: " ",
		1406: " ",
		1407: " ",
		1408: "L",
	}
	for year, isLeapStr := range testMap {
		isLeap := isLeapStr[0] == 'L'
		is.AddMsg("mismatch isLeap, year=%v", year).Equal(IsLeap(year), isLeap)
	}
}

func TestToJd(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[lib.Date]int{
		// FIXME: uncomment after switching to 33-years algorithm
		// {0, 1, 1}: 1947955,
		// {100, 1, 1}: 1984479,
		// {200, 1, 1}: 2021004,
		// {300, 1, 1}: 2057528,
		// {400, 1, 1}: 2094052,
		// {400, 2, 1}: 2094083,

		{1394, 1, 1}:  2457103,
		{1394, 2, 1}:  2457134,
		{1394, 3, 1}:  2457165,
		{1394, 4, 1}:  2457196,
		{1394, 5, 1}:  2457227,
		{1394, 6, 1}:  2457258,
		{1394, 7, 1}:  2457289,
		{1394, 8, 1}:  2457319,
		{1394, 9, 1}:  2457349,
		{1394, 10, 1}: 2457379,
		{1394, 11, 1}: 2457409,
		{1394, 12, 1}: 2457439,
		{1395, 1, 1}:  2457468,
		{1395, 2, 1}:  2457499,
		{1395, 3, 1}:  2457530,
		{1395, 4, 1}:  2457561,
		{1395, 5, 1}:  2457592,
		{1395, 6, 1}:  2457623,
		{1395, 7, 1}:  2457654,
		{1395, 8, 1}:  2457684,
		{1395, 9, 1}:  2457714,
		{1395, 10, 1}: 2457744,
		{1395, 11, 1}: 2457774,
		{1395, 12, 1}: 2457804,
		{1396, 1, 1}:  2457834,
		{1396, 2, 1}:  2457865,
		{1396, 3, 1}:  2457896,
		{1396, 4, 1}:  2457927,
		{1396, 5, 1}:  2457958,
		{1396, 6, 1}:  2457989,
		{1396, 7, 1}:  2458020,
		{1396, 8, 1}:  2458050,
		{1396, 9, 1}:  2458080,
		{1396, 10, 1}: 2458110,
		{1396, 11, 1}: 2458140,
		{1396, 12, 1}: 2458170,
	}
	for date, jd := range testMap {
		is.AddMsg("mismatch jd, date=%v, jd=%v", date, jd).Equal(ToJd(date), jd)
	}
}

func TestConvert(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	startYear := 1350
	endYear := 1450
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
