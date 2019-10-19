package main

import "fmt"

import (
	lib "github.com/ilius/libgostarcal"
	. "github.com/ilius/libgostarcal/cal_types/ethiopian"
)

func main() {
	PrintIsLeap(1990, 2030)
	// PrintToJd(2010, 2012)
	// PrintGetMonthLen(2015, 2018, 12)
}

func PrintIsLeap(startYear int, endYear int) {
	for year := startYear; year < endYear; year++ {
		fmt.Printf(
			"\t\t%v: %v,\n",
			year,
			IsLeap(year),
		)
	}
}

func PrintToJd(startYear int, endYear int) {
	var date *lib.Date
	var jd int
	for year := startYear; year < endYear; year++ {
		for month := uint8(1); month <= 12; month++ {
			date = lib.NewDate(year, month, 1)
			jd = ToJd(date)
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
				GetMonthLen(year, month),
			)
		}
	}
}
