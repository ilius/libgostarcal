package main

import "fmt"

import (
	lib "github.com/ilius/libgostarcal"
	. "github.com/ilius/libgostarcal/cal_types/gregorian"
)

func main() {
	PrintIsLeap(-50, 51)
	// PrintToJd(-50, 51, 1)
}

func PrintIsLeap(startYear int, endYear int) {
	for year := startYear; year < endYear; year++ {
		isLeapStr := " "
		if IsLeap(year) {
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
			jd = ToJd(date)
			fmt.Printf(
				"\t\tlib.%v: %v,\n",
				date.Repr(),
				jd,
			)
		}
	}
}
