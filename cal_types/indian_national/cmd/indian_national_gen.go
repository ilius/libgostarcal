package main

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
	"github.com/ilius/libgostarcal/cal_types/indian_national"
)

var calType = indian_national.New()

func main() {
	PrintIsLeap(1920, 1950)
	// PrintToJd(1950, 1952)
	// PrintGetMonthLen(1936, 1939, 12)
}

func PrintIsLeap(startYear int, endYear int) {
	for year := startYear; year < endYear; year++ {
		fmt.Printf(
			"\t\t%v: %v,\n",
			year,
			calType.IsLeap(year),
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
