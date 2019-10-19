package hijri

import (
	"testing"

	"github.com/ilius/is"

	lib "github.com/ilius/libgostarcal"
)

func TestIsLeap(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[int]bool{
		1410: false,
		1411: false,
		1412: true,
		1413: false,
		1414: false,
		1415: true,
		1416: false,
		1417: true,
		1418: false,
		1419: false,
		1420: true,
		1421: false,
		1422: false,
		1423: true,
		1424: false,
		1425: false,
		1426: true,
		1427: false,
		1428: true,
		1429: false,
		1430: false,
		1431: true,
		1432: false,
		1433: false,
		1434: true,
		1435: false,
		1436: true,
		1437: false,
		1438: false,
		1439: true,
		1440: false,
		1441: false,
		1442: true,
		1443: false,
		1444: false,
		1445: true,
		1446: false,
		1447: true,
		1448: false,
		1449: false,
	}
	for year, isLeap := range testMap {
		is.AddMsg("mismatch isLeap, year=%v", year).Equal(IsLeap(year), isLeap)
	}
}

func TestToJd(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[*lib.Date]int{
		lib.NewDate(1436, 1, 1):  2456957,
		lib.NewDate(1436, 2, 1):  2456987,
		lib.NewDate(1436, 3, 1):  2457016,
		lib.NewDate(1436, 4, 1):  2457046,
		lib.NewDate(1436, 5, 1):  2457075,
		lib.NewDate(1436, 6, 1):  2457105,
		lib.NewDate(1436, 7, 1):  2457134,
		lib.NewDate(1436, 8, 1):  2457164,
		lib.NewDate(1436, 9, 1):  2457193,
		lib.NewDate(1436, 10, 1): 2457223,
		lib.NewDate(1436, 11, 1): 2457252,
		lib.NewDate(1436, 12, 1): 2457282,
		lib.NewDate(1437, 1, 1):  2457312,
		lib.NewDate(1437, 2, 1):  2457342,
		lib.NewDate(1437, 3, 1):  2457371,
		lib.NewDate(1437, 4, 1):  2457401,
		lib.NewDate(1437, 5, 1):  2457430,
		lib.NewDate(1437, 6, 1):  2457460,
		lib.NewDate(1437, 7, 1):  2457489,
		lib.NewDate(1437, 8, 1):  2457519,
		lib.NewDate(1437, 9, 1):  2457548,
		lib.NewDate(1437, 10, 1): 2457578,
		lib.NewDate(1437, 11, 1): 2457607,
		lib.NewDate(1437, 12, 1): 2457637,
		lib.NewDate(1438, 1, 1):  2457666,
		lib.NewDate(1438, 2, 1):  2457696,
		lib.NewDate(1438, 3, 1):  2457725,
		lib.NewDate(1438, 4, 1):  2457755,
		lib.NewDate(1438, 5, 1):  2457784,
		lib.NewDate(1438, 6, 1):  2457814,
		lib.NewDate(1438, 7, 1):  2457843,
		lib.NewDate(1438, 8, 1):  2457873,
		lib.NewDate(1438, 9, 1):  2457902,
		lib.NewDate(1438, 10, 1): 2457932,
		lib.NewDate(1438, 11, 1): 2457961,
		lib.NewDate(1438, 12, 1): 2457991,
		lib.NewDate(1439, 1, 1):  2458020,
		lib.NewDate(1439, 2, 1):  2458050,
		lib.NewDate(1439, 3, 1):  2458079,
		lib.NewDate(1439, 4, 1):  2458109,
		lib.NewDate(1439, 5, 1):  2458138,
		lib.NewDate(1439, 6, 1):  2458168,
		lib.NewDate(1439, 7, 1):  2458197,
		lib.NewDate(1439, 8, 1):  2458227,
		lib.NewDate(1439, 9, 1):  2458256,
		lib.NewDate(1439, 10, 1): 2458286,
		lib.NewDate(1439, 11, 1): 2458315,
		lib.NewDate(1439, 12, 1): 2458345,
	}
	for date, jd := range testMap {
		is.AddMsg("mismatch jd, date=%v, jd=%v", date, jd).Equal(ToJd(date), jd)
	}
}

func TestMonthLen(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[[2]int]uint8{
		{1436, 1}:  30,
		{1436, 2}:  29,
		{1436, 3}:  30,
		{1436, 4}:  29,
		{1436, 5}:  30,
		{1436, 6}:  29,
		{1436, 7}:  30,
		{1436, 8}:  29,
		{1436, 9}:  30,
		{1436, 10}: 29,
		{1436, 11}: 30,
		{1436, 12}: 30,
		{1437, 1}:  30,
		{1437, 2}:  29,
		{1437, 3}:  30,
		{1437, 4}:  29,
		{1437, 5}:  30,
		{1437, 6}:  29,
		{1437, 7}:  30,
		{1437, 8}:  29,
		{1437, 9}:  30,
		{1437, 10}: 29,
		{1437, 11}: 30,
		{1437, 12}: 29,
		{1438, 1}:  30,
		{1438, 2}:  29,
		{1438, 3}:  30,
		{1438, 4}:  29,
		{1438, 5}:  30,
		{1438, 6}:  29,
		{1438, 7}:  30,
		{1438, 8}:  29,
		{1438, 9}:  30,
		{1438, 10}: 29,
		{1438, 11}: 30,
		{1438, 12}: 29,
		{1439, 1}:  30,
		{1439, 2}:  29,
		{1439, 3}:  30,
		{1439, 4}:  29,
		{1439, 5}:  30,
		{1439, 6}:  29,
		{1439, 7}:  30,
		{1439, 8}:  29,
		{1439, 9}:  30,
		{1439, 10}: 29,
		{1439, 11}: 30,
		{1439, 12}: 30,
	}
	for ym, mLen := range testMap {
		is.AddMsg("mismatch month length, ym=%v", ym).Equal(
			GetMonthLen(ym[0], uint8(ym[1])),
			mLen,
		)
	}
}

func TestConvert(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	startYear := 1390
	endYear := 1480
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
