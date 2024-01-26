// -*- coding: utf-8 -*-
//
// Copyright (C) Saeed Rasooli <saeed.gnu@gmail.com>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License along
// with this program. If not, see <https://www.gnu.org/licenses/agpl.txt>.

package hijri

import (
	"encoding/json"
	"math"

	lib "github.com/ilius/libgostarcal"

	. "github.com/ilius/libgostarcal/utils"
)

// ########## Types ##########

type MonthData struct {
	Version   [2]int  `json:"version"`
	StartDate [3]int  `json:"startDate"`
	StartJd   int     `json:"startJd"`
	MonthLen  [][]int `json:"monthLen"`
	ExpJd     int     `json:"expJd"`

	MonthLenByYm map[int]int `json:"-"`
	EndJd        int         `json:"-"`
}

// ###### Common Globals #######

var useMonthData = true

const (
	Name  = "hijri"
	Desc  = "Hijri(Islamic)"
	Epoch = 1948440

	MinMonthLen uint8 = 29
	MaxMonthLen uint8 = 30

	AvgYearLen = 354.3666 // FIXME
)

var MonthNames = []string{
	"Muharram", "Safar", "Rabia' 1",
	"Rabia' 2", "Jumada 1", "Jumada 2",
	"Rajab", "Sha'aban", "Ramadan",
	"Shawwal", "Dhu'l Qidah", "Dhu'l Hijjah",
}

var MonthNamesAb = []string{
	"Moh", "Saf", "Rb1",
	"Rb2", "Jm1", "Jm2",
	"Raj", "Shb", "Ram",
	"Shw", "DhQ", "DhH",
}

// ###### Other Globals  #######

const monthDataJSON = `{
	"version": [1443, 8],
	"startDate": [1426, 2, 1],
	"startJd": 2453442,
	"monthLen": [
		[1426,  0, 29, 30, 29, 30, 30, 30, 30, 29, 30, 29, 29],
		[1427, 30, 29, 29, 30, 29, 30, 30, 30, 30, 29, 29, 30],
		[1428, 29, 30, 29, 29, 29, 30, 30, 29, 30, 30, 30, 29],
		[1429, 30, 29, 30, 29, 29, 29, 30, 30, 29, 30, 30, 29],
		[1430, 30, 30, 29, 29, 30, 29, 30, 29, 29, 30, 30, 29],
		[1431, 30, 30, 29, 30, 29, 30, 29, 30, 29, 29, 30, 29],
		[1432, 30, 30, 29, 30, 30, 30, 29, 29, 30, 29, 30, 29],
		[1433, 29, 30, 29, 30, 30, 30, 29, 30, 29, 30, 29, 30],
		[1434, 29, 29, 30, 29, 30, 30, 29, 30, 30, 29, 30, 29],
		[1435, 29, 30, 29, 30, 29, 30, 29, 30, 30, 30, 29, 30],
		[1436, 29, 30, 29, 29, 30, 29, 30, 29, 30, 29, 30, 30],
		[1437, 29, 30, 30, 29, 30, 29, 29, 30, 29, 29, 30, 30],
		[1438, 29, 30, 30, 30, 29, 30, 29, 29, 30, 29, 29, 30],
		[1439, 29, 30, 30, 30, 30, 29, 30, 29, 29, 30, 29, 29],
		[1440, 30, 29, 30, 30, 30, 29, 30, 30, 29, 29, 30, 29],
		[1441, 29, 30, 29, 30, 30, 29, 30, 30, 29, 30, 29, 30],
		[1442, 29, 29, 30, 29, 30, 29, 30, 30, 29, 30, 30, 29],
		[1443, 29, 30, 30, 29, 29, 30, 29, 30]
	],
	"expJd": 2459660
}`

var monthData *MonthData

// ######## MonthData methods ########

func (mdata *MonthData) Load() {
	monthLenByYear := map[int][]int{}
	for _, row := range mdata.MonthLen {
		monthLenByYear[row[0]] = row[1:]
	}
	mdata.setMonthLenByYear(monthLenByYear)
}

func (mdata *MonthData) setMonthLenByYear(monthLenByYear map[int][]int) {
	endJd := mdata.StartJd
	monthLenByYm := map[int]int{}
	for year, lengthList := range monthLenByYear {
		for mm, length := range lengthList {
			// mm is month - 1
			if length <= 0 {
				continue
			}
			monthLenByYm[year*12+mm] = length
			endJd += length
		}
	}
	mdata.MonthLenByYm = monthLenByYm
	mdata.EndJd = endJd
}

func (mdata *MonthData) GetDateFromJd(jd int) *lib.Date {
	if !(mdata.EndJd >= jd && jd >= mdata.StartJd) {
		return nil
	}
	y := mdata.StartDate[0]
	m := mdata.StartDate[1]
	d := mdata.StartDate[2]
	ym := y*12 + m - 1
	startJd := mdata.StartJd
	for jd > startJd {
		monthLen := mdata.MonthLenByYm[ym]
		jdm0 := jd - monthLen

		if jdm0 <= startJd-d {
			d = d + jd - startJd
			break
		}

		if startJd-d < jdm0 && jdm0 <= startJd {
			ym += 1
			d = d + jd - startJd - monthLen
			break
		}

		// assert(jdm0 > startJd)
		ym += 1
		jd -= monthLen
	}
	year, mm := Divmod(ym, 12)
	return lib.NewDate(
		year,
		uint8(mm+1),
		uint8(d),
	)
}

func (mdata *MonthData) GetJdFromDate(date *lib.Date) (int, bool) {
	year := date.Year
	ym := year*12 + int(date.Month) - 1
	_, ok := mdata.MonthLenByYm[ym-1]
	if !ok {
		return 0, false
	}
	ym0 := mdata.StartDate[0]*12 + mdata.StartDate[1] - 1
	jd := mdata.StartJd
	for ymi := ym0; ymi < ym; ymi++ {
		plus, ok := mdata.MonthLenByYm[ymi]
		if !ok {
			panic("mdata.MonthLenByYm[ymi]")
		}
		jd += plus
	}
	return jd + int(date.Day) - 1, true
}

// #############################

func New() *calTypeImp {
	return &calTypeImp{}
}

type calTypeImp struct{}

func (*calTypeImp) Name() string {
	return Name
}

func (*calTypeImp) Desc() string {
	return Desc
}

func (*calTypeImp) Epoch() int {
	return Epoch
}

func (*calTypeImp) MinMonthLen() uint8 {
	return MinMonthLen
}

func (*calTypeImp) MaxMonthLen() uint8 {
	return MaxMonthLen
}

func (*calTypeImp) AvgYearLen() float64 {
	return AvgYearLen
}

func (*calTypeImp) MonthNames() []string {
	return MonthNames
}

func (*calTypeImp) MonthNamesAb() []string {
	return MonthNamesAb
}

func SetUseMonthData(use bool) {
	if !use {
		useMonthData = false
		return
	}
	if monthData == nil {
		monthData = &MonthData{}
		err := json.Unmarshal([]byte(monthDataJSON), monthData)
		if err != nil {
			panic(err)
		}
		monthData.Load()
	}
	useMonthData = true
}

func (*calTypeImp) IsLeap(year int) bool {
	return Mod(year*11+14, 30) < 11
}

func (*calTypeImp) ToJd(date *lib.Date) int {
	if useMonthData { // and HijriAlg==0
		jd, ok := monthData.GetJdFromDate(date)
		if ok {
			return jd
		}
	}
	return (int(date.Day) +
		int(math.Ceil(29.5*float64(date.Month-1))) +
		(date.Year-1)*354 +
		Div(11*date.Year+3, 30) +
		Epoch)
}

func (ct *calTypeImp) JdTo(jd int) *lib.Date {
	// jdf := jd + 0.5
	if useMonthData { // && HijriAlg==0
		date := monthData.GetDateFromJd(jd)
		if date != nil {
			return date
		}
	}
	year := Div(30*(jd-1-Epoch)+10646, 10631)
	month := uint8(IntMin(
		12,
		int(math.Ceil(
			(float64(jd)+0.5-float64(ct.ToJd(lib.NewDate(year, 1, 1))))/29.5,
		)),
	))
	day := uint8(jd - ct.ToJd(lib.NewDate(year, month, 1)) + 1)
	return lib.NewDate(year, month, day)
}

func (ct *calTypeImp) GetMonthLen(year int, month uint8) uint8 {
	if useMonthData { // && HijriAlg==0
		if month == 12 {
			return uint8(
				ct.ToJd(lib.NewDate(year+1, 1, 1)) - ct.ToJd(lib.NewDate(year, 12, 1)),
			)
		}
		return uint8(
			ct.ToJd(lib.NewDate(year, month+1, 1)) - ct.ToJd(lib.NewDate(year, month, 1)),
		)
	}
	if month%2 == 1 { // safe %
		return 30
	}
	if month == 12 && ct.IsLeap(year) {
		return 30
	}
	return 29
}
