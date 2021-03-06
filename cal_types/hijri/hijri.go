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
	"math"

	lib "github.com/ilius/libgostarcal"

	"github.com/ilius/libgostarcal/cal_types"
	. "github.com/ilius/libgostarcal/utils"
)

// ###### Common Globals #######

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

// #############################

func init() {
	cal_types.RegisterCalType(
		Name,
		Desc,
		Epoch,
		MinMonthLen,
		MaxMonthLen,
		AvgYearLen,
		MonthNames,
		MonthNamesAb,
		IsLeap,
		ToJd,
		JdTo,
		GetMonthLen,
	)
}

func IsLeap(year int) bool {
	return Mod(year*11+14, 30) < 11
}

func ToJd(date *lib.Date) int {
	return (int(date.Day) +
		int(math.Ceil(29.5*float64(date.Month-1))) +
		(date.Year-1)*354 +
		Div(11*date.Year+3, 30) +
		Epoch)
}

func JdTo(jd int) *lib.Date {
	// jdf := jd + 0.5
	year := Div(30*(jd-1-Epoch)+10646, 10631)
	month := uint8(IntMin(
		12,
		int(math.Ceil(
			(float64(jd)+0.5-float64(ToJd(lib.NewDate(year, 1, 1))))/29.5,
		)),
	))
	day := uint8(jd - ToJd(lib.NewDate(year, month, 1)) + 1)
	return lib.NewDate(year, month, day)
}

func GetMonthLen(year int, month uint8) uint8 {
	if month%2 == 1 { // safe %
		return 30
	}
	if month == 12 && IsLeap(year) {
		return 30
	}
	return 29
}
