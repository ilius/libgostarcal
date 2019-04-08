package indian_national

import "fmt"

import lib "github.com/ilius/libgostarcal"

func PrintIsLeap(startYear int, endYear int) {
	for year := startYear; year < endYear; year++ {
		fmt.Printf(
			"        %v: %v,\n",
			year,
			IsLeap(year),
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
