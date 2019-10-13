package main

import "fmt"

import (
	lib "github.com/ilius/libgostarcal"
	. "github.com/ilius/libgostarcal/cal_types/gregorian"
)

func main() {
	PrintIsLeap(1990, 2030)
}

func PrintIsLeap(startYear int, endYear int) {
	for year := startYear; year < endYear; year++ {
		isLeapStr := " "
		if IsLeap(year) {
			isLeapStr = "L"
		}
		fmt.Printf(
			"        %v: \"%v\",\n",
			year,
			isLeapStr,
		)
	}
}

func PrintToJd(startYear int, endYear int) {
	var date lib.Date
	var jd int
	for year := startYear; year < endYear; year++ {
		for month := uint8(1); month <= 12; month++ {
			date = lib.Date{year, month, 1}
			jd = ToJd(date)
			fmt.Printf(
				"        %v: %v,\n",
				date.Repr(),
				jd,
			)
		}
	}
}