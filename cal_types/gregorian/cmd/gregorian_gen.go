package main

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
	"github.com/ilius/libgostarcal/cal_types/gregorian"
)

var calType = gregorian.New()

func main() {
	PrintIsLeap(-50, 51)
	// PrintToJd(-50, 51, 1)
	// PrintGetMonthLen(2015, 2018, 12)
}

func PrintIsLeap(startYear int, endYear int) {
	for year := startYear; year < endYear; year++ {
		isLeapStr := " "
		if calType.IsLeap(year) {
			isLeapStr = "L"
		}
		fmt.Printf(
			"\t\t%v: \"%v\",\n",
			year,
			isLeapStr,
		)
	}
}

func PrintToJd(startYear int, endYear int, maxMonth uint8) {
	var date *lib.Date
	var jd int
	for year := startYear; year < endYear; year++ {
		for month := uint8(1); month <= maxMonth; month++ {
			date = lib.NewDate(year, month, 1)
			jd = calType.ToJd(date)
			fmt.Printf(
				"\t\tlib.%v: %v,\n",
				date.Repr(),
				jd,
			)
		}
	}
}

func PrintGetMonthLen(startYear int, endYear int, maxMonth uint8) {
	for year := startYear; year < endYear; year++ {
		for month := uint8(1); month <= maxMonth; month++ {
			fmt.Printf(
				"\t\t{%v, %v}: %v,\n",
				year, month,
				calType.GetMonthLen(year, month),
			)
		}
	}
}
