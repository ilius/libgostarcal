// -*- coding: utf-8 -*-
//
// Copyright (C) Saeed Rasooli <saeed.gnu@gmail.com>
// Copyright (C) 2007 Mehdi Bayazee <Bayazee@Gmail.com>
// Copyright (C) 2001 Roozbeh Pournader <roozbeh@sharif.edu>
// Copyright (C) 2001 Mohammad Toossi <mohammad@bamdad.org>
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

// Iranian (Jalali) calendar:
// http://en.wikipedia.org/wiki/Iranian_calendar

package jalali

import (
	"log"

	lib "github.com/ilius/libgostarcal"
	. "github.com/ilius/libgostarcal/utils"
)

// ###### Common Globals #######

const (
	Name  = "jalali"
	Desc  = "Jalali"
	Epoch = 1948321

	MinMonthLen uint8 = 29
	MaxMonthLen uint8 = 31

	AvgYearLen = 365.2425 // FIXME

	GREGORIAN_EPOCH = 1721426
)

var MonthNames = []string{
	"Farvardin", "Ordibehesht", "Khordad", "Teer", "Mordad", "Shahrivar",
	"Mehr", "Aban", "Azar", "Dey", "Bahman", "Esfand",
}

var MonthNamesAb = []string{
	"Far", "Ord", "Khr", "Tir", "Mor", "Shr",
	"Meh", "Abn", "Azr", "Dey", "Bah", "Esf",
}

// ###### Other Globals  #######

var (
	monthLen    = []uint8{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 30}
	monthLenSum = []int{0, 31, 62, 93, 124, 155, 186, 216, 246, 276, 306, 336, 366}
)

var alg2820 bool = false

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

// SetAlgorithm2820: enable 2820-year algorithm by passing true
func SetAlgorithm2820(useAlg2820 bool) {
	log.Printf("jalali.go: SetAlgorithm2820(%v)\n", useAlg2820)
	alg2820 = useAlg2820
}

// IsLeap: return true if year is leap, false otherwise
// Normal: esfand = 29 days
// Leap: esfand = 30 days
func (*calTypeImp) IsLeap(year int) bool {
	if alg2820 {
		// using 2820-years algorithm
		/* if we want to remove Year Zero:
		if year < 1 {
			year++
		}
		*/
		return Mod((Mod(year-474, 2820))*682, 2816) < 682
	}
	jy := year - 979
	jyd, jym := Divmod(jy, 33)
	jyd2, jym2 := Divmod(jy+1, 33)
	return 1 == (jyd2-jyd)*8+(jym2+3)/4-(jym+3)/4
}

// ToJd: calculate Julian day from Jalali date
func (*calTypeImp) ToJd(date *lib.Date) int {
	if alg2820 {
		// using 2820-years algorithm
		epbase := date.Year - 474
		/* if we want to remove Year Zero:
		if date.Year < 0 {
			epbase += 1
		}
		*/
		epbase_d, epbase_m := Divmod(epbase, 2820)
		epyear := 474 + epbase_m
		mm := int(date.Month - 1)
		return int(date.Day) +
			mm*30 + IntMin(6, mm) +
			Div(epyear*682-110, 2816) +
			(epyear-1)*365 +
			epbase_d*1029983 +
			Epoch - 1
	}
	jy := date.Year - 979
	jyd, jym := Divmod(jy, 33)
	return 365*jy +
		jyd*8 +
		Div(jym+3, 4) +
		monthLenSum[date.Month-1] +
		int(date.Day) - 1 +
		584101 +
		GREGORIAN_EPOCH
}

func getMonthDayFromYdays(yday int) (uint8, uint8) {
	month := uint8(BisectLeft(monthLenSum, yday))
	day := uint8(yday - monthLenSum[month-1])
	return month, day
}

// JdTo: calculate Jalali date from Julian day
func (ct *calTypeImp) JdTo(jd int) *lib.Date {
	if alg2820 {
		// using 2820-years algorithm
		deltaDays := jd - ct.ToJd(lib.NewDate(475, 1, 1))
		cycle, cyear := Divmod(deltaDays, 1029983)
		var ycycle int
		if cyear == 1029982 {
			ycycle = 2820
		} else {
			aux1, aux2 := Divmod(cyear, 366)
			// cyear >= 0, aux2 >= 0
			ycycle = Div(2134*aux1+2816*aux2+2815, 1028522) + cyear/366 + 1
		}
		year := 2820*cycle + ycycle + 474
		/* if we want to remove Year Zero:
		if year <= 0 {
			year--
		}
		*/
		yday := jd - ct.ToJd(lib.NewDate(year, 1, 1)) + 1
		month, day := getMonthDayFromYdays(yday)
		return lib.NewDate(year, month, day)
	}
	jdays := jd - GREGORIAN_EPOCH - 584101
	// -(1600*365 + 1600//4 - 1600//100 + 1600//400) + 365-79+1 == -584101
	j_np, jdays := Divmod(jdays, 12053)

	yearFact, jdays := Divmod(jdays, 1461)
	year := 979 + 33*j_np + 4*yearFact

	if jdays >= 366 {
		var yearPlus int
		yearPlus, jdays = Divmod(jdays-1, 365)
		year += yearPlus
	}
	yday := jdays + 1
	month, day := getMonthDayFromYdays(yday)
	return lib.NewDate(year, month, day)
}

func (ct *calTypeImp) GetMonthLen(year int, month uint8) uint8 {
	if month == 12 {
		if ct.IsLeap(year) {
			return 30
		}
		return 29
	}
	return monthLen[month-1]
}
