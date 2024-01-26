package main

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
	"github.com/ilius/libgostarcal/cal_types/hijri"
)

var calType = hijri.New()

func main() {
	PrintIsLeap(1410, 1450)
	// PrintToJd(1450, 1452)
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
