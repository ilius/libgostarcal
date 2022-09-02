package interval

import (
	"math/rand"
	"testing"
	"time"

	"github.com/ilius/is/v2"
)

func ShuffleIntervals(a []*Interval) {
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func TestIntervalListByNumList(t *testing.T) {
	is := is.New(t)
	nums := []int64{1, 2, 3, 4, 5, 7, 9, 10, 14, 16, 17, 18, 19, 21, 22, 23, 24}
	intervals := IntervalListByNumList(nums, 3)
	is.Equal(intervals.String(), "1-5] 7 9 10 14 16-19] 21-24]")
}

func TestIntervalListNormalize(t *testing.T) {
	test := func(t *testing.T, input string, normalized string) {
		is := is.New(t).MsgSep(", ").Msg("input=%#v, normalized=%#v", input, normalized)
		t.Helper()

		testList, testErr := ParseIntervalList(input)
		is.NotErr(testErr)

		answerList, answerErr := ParseIntervalList(normalized)
		is.NotErr(answerErr)

		ShuffleIntervals(testList)

		testList, testErr = testList.Normalize()
		is.NotErr(testErr)

		testList = testList.Humanize()
		is.PrependMsg(
			"mismatch result, testList=%#v, answerList=%#v",
			testList.String(),
			answerList.String(),
		).Equal(testList.String(), answerList.String())
		t.Log(input, "=>", answerList)
	}

	test(t, "10-20 20", "10-20 20")
	test(t, "10-20]", "10-20 20")
	test(t, "10-20 20 20", "10-20 20")
	test(t, "10-20] 20-30]", "10-30 30")
	test(t, "10-20 20-30", "10-30")
	test(t, "10-20 20-30 30-40", "10-40")
	test(t, "10-20 20-30 25-40", "10-40")
	test(t, "1-10 14 2-5 9-13 16-18 17-20 15-16 25-30", "1-13 14 15-20 25-30")
	test(t, "60-70 0-40 10-50 20-30 80-90 70-80 85-100 110 55", "0-50 55 60-100 110")
	test(t, "-(70-60) -(60-55) -(56-50)", "-(70-50)")
	test(t, "-(70-60) -(60-55) -(56-50])", "-(70-50) -50")
}

func getIntersectionString(t *testing.T, list1Str string, list2Str string) string {
	list1, err1 := ParseIntervalList(list1Str)
	list2, err2 := ParseIntervalList(list2Str)
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err1)
	}
	ShuffleIntervals(list1)
	ShuffleIntervals(list2)

	result, err := list1.Intersection(list2)
	if err != nil {
		t.Error(err)
		return ""
	}
	result = result.Humanize()
	return result.String()
}

func TestIntervalListIntersection(t *testing.T) {
	type p [2]string

	testMap := map[[2]string]string{
		p{
			"0-20",
			"10-30",
		}: "10-20",

		p{
			"10-30 40-50 60-80",
			"25-45",
		}: "25-30 40-45",

		p{
			"10-30 40-50 60-80",
			"25-45 50-60",
		}: "25-30 40-45",

		p{
			"10-30 40-50 60-80",
			"25-45 50-60 60",
		}: "25-30 40-45 60",

		p{
			"10-30 40-50 60-80",
			"25-45 48-70 60",
		}: "25-30 40-45 48-50 60-70",

		p{
			"10-30 40-50 60-80",
			"25-45 48-70",
		}: "25-30 40-45 48-50 60-70",

		p{
			"0-10 20-30 40-50 60-70",
			"1-2 6-7 11-12 16-17 21-22 26-27 27",
		}: "1-2 6-7 21-22 26-27 27",

		/*
		   p{
		       "",
		       "",
		   }:  "",
		*/

	}

	for testPair, answerStr := range testMap {
		resultStr := getIntersectionString(t, testPair[0], testPair[1])
		if resultStr != answerStr {
			t.Error("test failed:")
			t.Error("resultStr =", resultStr)
			t.Error("answerStr =", answerStr)
		}
	}
}
