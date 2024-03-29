package main

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
	"github.com/ilius/libgostarcal/cal_types/jalali"
)

var calType = jalali.New()

func main() {
	PrintIsLeap(-100, 11)
	PrintIsLeap(1360, 1409)
	// PrintToJd(1380, 1390)
	// PrintGetMonthLen(1394, 1397, 12)
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

func PrintToJd(startYear int, endYear int) {
	var date *lib.Date
	var jd int
	for year := startYear; year < endYear; year++ {
		for month := uint8(1); month <= 12; month++ {
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
