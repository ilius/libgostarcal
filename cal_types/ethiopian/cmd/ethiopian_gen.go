package main

import "fmt"

import (
	lib "github.com/ilius/libgostarcal"
	. "github.com/ilius/libgostarcal/cal_types/ethiopian"
)

func main() {
	PrintIsLeap(1990, 2030)
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
				"\t\t%v: %v,\n",
				date.Repr(),
				jd,
			)
		}
	}
}