package jalali

import (
	"sort"
	"testing"

	"github.com/ilius/is"

	lib "github.com/ilius/libgostarcal"
)

func lastMonthLenByConvert(year int) int {
	return ToJd(lib.NewDate(year+1, 1, 1)) - ToJd(lib.NewDate(year, 12, 1))
}

func TestIsLeap(t *testing.T) {
	defer SetAlgorithm2820(alg2820)
	is := is.New(t).Lax().MsgSep(", ")
	// the values are 2-letter strings
	// the first letter is for 33-year algorithm (which is not implemented yet)
	// the second lettee is for 2820 algorithm (currently used)
	testMap := map[int]string{
		10:   "  ",
		9:    "L ",
		8:    " L",
		7:    "  ",
		6:    "  ",
		5:    "L ",
		4:    " L",
		3:    "  ",
		2:    "  ",
		1:    "L ",
		0:    " L",
		-1:   "  ",
		-2:   "  ",
		-3:   "L ",
		-4:   " L",
		-5:   "  ",
		-6:   "  ",
		-7:   "L ",
		-8:   " L",
		-9:   "  ",
		-10:  "  ",
		-11:  "L ",
		-12:  "  ",
		-13:  " L",
		-14:  "  ",
		-15:  "  ",
		-16:  "L ",
		-17:  " L",
		-18:  "  ",
		-19:  "  ",
		-20:  "L ",
		-21:  " L",
		-22:  "  ",
		-23:  "  ",
		-24:  "L ",
		-25:  " L",
		-26:  "  ",
		-27:  "  ",
		-28:  "L ",
		-29:  " L",
		-30:  "  ",
		-31:  "  ",
		-32:  "L ",
		-33:  " L",
		-34:  "  ",
		-35:  "  ",
		-36:  "L ",
		-37:  " L",
		-38:  "  ",
		-39:  "  ",
		-40:  "L ",
		-41:  "  ",
		-42:  " L",
		-43:  "  ",
		-44:  "L ",
		-45:  "  ",
		-46:  " L",
		-47:  "  ",
		-48:  "  ",
		-49:  "L ",
		-50:  " L",
		-51:  "  ",
		-52:  "  ",
		-53:  "L ",
		-54:  " L",
		-55:  "  ",
		-56:  "  ",
		-57:  "L ",
		-58:  " L",
		-59:  "  ",
		-60:  "  ",
		-61:  "L ",
		-62:  " L",
		-63:  "  ",
		-64:  "  ",
		-65:  "L ",
		-66:  " L",
		-67:  "  ",
		-68:  "  ",
		-69:  "L ",
		-70:  " L",
		-71:  "  ",
		-72:  "  ",
		-73:  "L ",
		-74:  "  ",
		-75:  " L",
		-76:  "  ",
		-77:  "L ",
		-78:  "  ",
		-79:  " L",
		-80:  "  ",
		-81:  "  ",
		-82:  "L ",
		-83:  " L",
		-84:  "  ",
		-85:  "  ",
		-86:  "L ",
		-87:  " L",
		-88:  "  ",
		-89:  "  ",
		-90:  "L ",
		-91:  " L",
		-92:  "  ",
		-93:  "  ",
		-94:  "L ",
		-95:  " L",
		-96:  "  ",
		-97:  "  ",
		-98:  "L ",
		-99:  " L",
		-100: "  ",

		1360: "  ",
		1361: "  ",
		1362: "LL",
		1363: "  ",
		1364: "  ",
		1365: "  ",
		1366: "LL",
		1367: "  ",
		1368: "  ",
		1369: "  ",
		1370: "LL",
		1371: "  ",
		1372: "  ",
		1373: "  ",
		1374: "  ",
		1375: "LL",
		1376: "  ",
		1377: "  ",
		1378: "  ",
		1379: "LL",
		1380: "  ",
		1381: "  ",
		1382: "  ",
		1383: "LL",
		1384: "  ",
		1385: "  ",
		1386: "  ",
		1387: "LL",
		1388: "  ",
		1389: "  ",
		1390: "  ",
		1391: "LL",
		1392: "  ",
		1393: "  ",
		1394: "  ",
		1395: "LL",
		1396: "  ",
		1397: "  ",
		1398: "  ",
		1399: "LL",
		1400: "  ",
		1401: "  ",
		1402: "  ",
		1404: " L", // FIXME: why mismatch
		1405: "  ",
		1406: "  ",
		1407: "  ",
		1408: "LL",
	}
	for algIndex, alg2820 := range []bool{false, true} {
		SetAlgorithm2820(alg2820)
		testList := make([][]interface{}, 0, len(testMap))
		for key, value := range testMap {
			testList = append(testList, []interface{}{key, value})
		}
		sort.Slice(testList, func(i, j int) bool {
			return testList[i][0].(int) < testList[j][0].(int)
		})
		for _, item := range testList {
			year := item[0].(int)
			isLeapStr := item[1].(string)
			isLeap := isLeapStr[algIndex] == 'L'
			is.AddMsg(
				"mismatch isLeap, year=%v, alg2820=%v",
				year,
				alg2820,
			).Equal(IsLeap(year), isLeap)
			lastMonthLen := lastMonthLenByConvert(year)
			lastMonthLenExpected := 29
			if isLeap {
				lastMonthLenExpected = 30
			}
			is.AddMsg(
				"year=%v (%v), alg2820=%v",
				year,
				string(isLeapStr[algIndex]),
				alg2820,
			).Equal(lastMonthLen, lastMonthLenExpected)
		}
	}
}

func TestToJd(t *testing.T) {
	defer SetAlgorithm2820(alg2820)
	is := is.New(t).Lax().MsgSep(", ")
	testMap := map[*lib.Date][2]int{
		lib.NewDate(0, 1, 1):   {1947955, 1947955},
		lib.NewDate(100, 1, 1): {1984479, 1984480}, // mismatch
		lib.NewDate(200, 1, 1): {2021004, 2021004},
		lib.NewDate(300, 1, 1): {2057528, 2057528},
		lib.NewDate(400, 1, 1): {2094052, 2094052},
		lib.NewDate(400, 2, 1): {2094083, 2094083},

		lib.NewDate(1394, 1, 1):  {2457103, 2457103},
		lib.NewDate(1394, 2, 1):  {2457134, 2457134},
		lib.NewDate(1394, 3, 1):  {2457165, 2457165},
		lib.NewDate(1394, 4, 1):  {2457196, 2457196},
		lib.NewDate(1394, 5, 1):  {2457227, 2457227},
		lib.NewDate(1394, 6, 1):  {2457258, 2457258},
		lib.NewDate(1394, 7, 1):  {2457289, 2457289},
		lib.NewDate(1394, 8, 1):  {2457319, 2457319},
		lib.NewDate(1394, 9, 1):  {2457349, 2457349},
		lib.NewDate(1394, 10, 1): {2457379, 2457379},
		lib.NewDate(1394, 11, 1): {2457409, 2457409},
		lib.NewDate(1394, 12, 1): {2457439, 2457439},
		lib.NewDate(1395, 1, 1):  {2457468, 2457468},
		lib.NewDate(1395, 2, 1):  {2457499, 2457499},
		lib.NewDate(1395, 3, 1):  {2457530, 2457530},
		lib.NewDate(1395, 4, 1):  {2457561, 2457561},
		lib.NewDate(1395, 5, 1):  {2457592, 2457592},
		lib.NewDate(1395, 6, 1):  {2457623, 2457623},
		lib.NewDate(1395, 7, 1):  {2457654, 2457654},
		lib.NewDate(1395, 8, 1):  {2457684, 2457684},
		lib.NewDate(1395, 9, 1):  {2457714, 2457714},
		lib.NewDate(1395, 10, 1): {2457744, 2457744},
		lib.NewDate(1395, 11, 1): {2457774, 2457774},
		lib.NewDate(1395, 12, 1): {2457804, 2457804},
		lib.NewDate(1396, 1, 1):  {2457834, 2457834},
		lib.NewDate(1396, 2, 1):  {2457865, 2457865},
		lib.NewDate(1396, 3, 1):  {2457896, 2457896},
		lib.NewDate(1396, 4, 1):  {2457927, 2457927},
		lib.NewDate(1396, 5, 1):  {2457958, 2457958},
		lib.NewDate(1396, 6, 1):  {2457989, 2457989},
		lib.NewDate(1396, 7, 1):  {2458020, 2458020},
		lib.NewDate(1396, 8, 1):  {2458050, 2458050},
		lib.NewDate(1396, 9, 1):  {2458080, 2458080},
		lib.NewDate(1396, 10, 1): {2458110, 2458110},
		lib.NewDate(1396, 11, 1): {2458140, 2458140},
		lib.NewDate(1396, 12, 1): {2458170, 2458170},
	}
	for algIndex, alg2820 := range []bool{false, true} {
		SetAlgorithm2820(alg2820)
		for date, jdByAlg := range testMap {
			jd := jdByAlg[algIndex]
			is.AddMsg(
				"mismatch jd, date=%v, jd=%v, alg2820=%v",
				date,
				jd,
				alg2820,
			).Equal(ToJd(date), jd)
		}
	}
}

func TestConvert(t *testing.T) {
	defer SetAlgorithm2820(alg2820)
	is := is.New(t).Lax().MsgSep(", ")
	startYear := 1350
	endYear := 1450
	for _, alg2820 := range []bool{false, true} {
		SetAlgorithm2820(alg2820)
		for year := startYear; year < endYear; year++ {
			for month := uint8(1); month <= 12; month++ {
				monthLen := GetMonthLen(year, month)
				for day := uint8(1); day <= monthLen; day++ {
					date := lib.NewDate(year, month, day)
					jd := ToJd(date)
					ndate := JdTo(jd)
					is.AddMsg(
						"jd=%v, date=%v, ndate=%v, alg2820=%v",
						jd,
						date,
						ndate,
						alg2820,
					).Equal(ndate, date)
				}
			}
		}
	}
}
