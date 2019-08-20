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

package gregorian

import (
	"time"

	lib "github.com/ilius/libgostarcal"
	"github.com/ilius/libgostarcal/cal_types"
)

// ###### Common Globals #######

const (
	Name  = "gregorian"
	Desc  = "Gregorian"
	Epoch = 1721426

	MinMonthLen uint8 = 29
	MaxMonthLen uint8 = 31

	AvgYearLen = 365.2425 // FIXME
)

var MonthNames = []string{
	"January", "February", "March", "April", "May", "June",
	"July", "August", "September", "October", "November", "December",
}

var MonthNamesAb = []string{
	"Jan", "Feb", "Mar", "Apr", "May", "Jun",
	"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
}

// ###### Other Globals  #######

// const J0001 = Epoch
const J1970 = 2440588

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
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func ToJd(date lib.Date) int {
	t := time.Date(
		date.Year,
		time.Month(date.Month),
		int(date.Day),
		0, 0, 0,
		0,
		time.UTC,
	)
	return J1970 + int(t.Unix()/86400)
}

func JdTo(jd int) lib.Date {
	t := time.Unix(
		int64(86400*(jd-J1970)),
		0,
	)
	return lib.Date{
		t.Year(),
		uint8(t.Month()),
		uint8(t.Day()),
	}
}

func GetMonthLen(year int, month uint8) uint8 {
	if month == 12 {
		return uint8(ToJd(lib.Date{year + 1, 1, 1}) - ToJd(lib.Date{year, 12, 1}))
	} else {
		return uint8(ToJd(lib.Date{year, month + 1, 1}) - ToJd(lib.Date{year, month, 1}))
	}
}
