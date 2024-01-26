package indian_national

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
		1920: false,
		1921: false,
		1922: true,
		1923: false,
		1924: false,
		1925: false,
		1926: true,
		1927: false,
		1928: false,
		1929: false,
		1930: true,
		1931: false,
		1932: false,
		1933: false,
		1934: true,
		1935: false,
		1936: false,
		1937: false,
		1938: true,
		1939: false,
		1940: false,
		1941: false,
		1942: true,
		1943: false,
		1944: false,
		1945: false,
		1946: true,
		1947: false,
		1948: false,
		1949: false,
	}
	for year, isLeap := range testMap {
		is.AddMsg("mismatch isLeap, year=%v", year).Equal(calType.IsLeap(year), isLeap)
	}
}

func TestToJd(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[*lib.Date]int{
		lib.NewDate(1936, 1, 1):  2456739,
		lib.NewDate(1936, 2, 1):  2456769,
		lib.NewDate(1936, 3, 1):  2456800,
		lib.NewDate(1936, 4, 1):  2456831,
		lib.NewDate(1936, 5, 1):  2456862,
		lib.NewDate(1936, 6, 1):  2456893,
		lib.NewDate(1936, 7, 1):  2456924,
		lib.NewDate(1936, 8, 1):  2456954,
		lib.NewDate(1936, 9, 1):  2456984,
		lib.NewDate(1936, 10, 1): 2457014,
		lib.NewDate(1936, 11, 1): 2457044,
		lib.NewDate(1936, 12, 1): 2457074,
		lib.NewDate(1937, 1, 1):  2457104,
		lib.NewDate(1937, 2, 1):  2457134,
		lib.NewDate(1937, 3, 1):  2457165,
		lib.NewDate(1937, 4, 1):  2457196,
		lib.NewDate(1937, 5, 1):  2457227,
		lib.NewDate(1937, 6, 1):  2457258,
		lib.NewDate(1937, 7, 1):  2457289,
		lib.NewDate(1937, 8, 1):  2457319,
		lib.NewDate(1937, 9, 1):  2457349,
		lib.NewDate(1937, 10, 1): 2457379,
		lib.NewDate(1937, 11, 1): 2457409,
		lib.NewDate(1937, 12, 1): 2457439,
		lib.NewDate(1938, 1, 1):  2457469,
		lib.NewDate(1938, 2, 1):  2457500,
		lib.NewDate(1938, 3, 1):  2457531,
		lib.NewDate(1938, 4, 1):  2457562,
		lib.NewDate(1938, 5, 1):  2457593,
		lib.NewDate(1938, 6, 1):  2457624,
		lib.NewDate(1938, 7, 1):  2457655,
		lib.NewDate(1938, 8, 1):  2457685,
		lib.NewDate(1938, 9, 1):  2457715,
		lib.NewDate(1938, 10, 1): 2457745,
		lib.NewDate(1938, 11, 1): 2457775,
		lib.NewDate(1938, 12, 1): 2457805,
	}
	for date, jd := range testMap {
		is.AddMsg("mismatch jd, date=%v, jd=%v", date, jd).Equal(calType.ToJd(date), jd)
	}
}

func TestGetMonthLen(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[[2]int]int{
		{1936, 1}:  30,
		{1936, 2}:  31,
		{1936, 3}:  31,
		{1936, 4}:  31,
		{1936, 5}:  31,
		{1936, 6}:  31,
		{1936, 7}:  30,
		{1936, 8}:  30,
		{1936, 9}:  30,
		{1936, 10}: 30,
		{1936, 11}: 30,
		{1936, 12}: 30,
		{1937, 1}:  30,
		{1937, 2}:  31,
		{1937, 3}:  31,
		{1937, 4}:  31,
		{1937, 5}:  31,
		{1937, 6}:  31,
		{1937, 7}:  30,
		{1937, 8}:  30,
		{1937, 9}:  30,
		{1937, 10}: 30,
		{1937, 11}: 30,
		{1937, 12}: 30,
		{1938, 1}:  31,
		{1938, 2}:  31,
		{1938, 3}:  31,
		{1938, 4}:  31,
		{1938, 5}:  31,
		{1938, 6}:  31,
		{1938, 7}:  30,
		{1938, 8}:  30,
		{1938, 9}:  30,
		{1938, 10}: 30,
		{1938, 11}: 30,
		{1938, 12}: 30,
	}
	for ym, monthLen := range testMap {
		year := ym[0]
		month := uint8(ym[1])
		is.AddMsg("ym={%v, %v}", year, month).Equal(calType.GetMonthLen(year, month), monthLen)
	}
}

func TestConvert(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	startYear := 1920
	endYear := 2950
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
