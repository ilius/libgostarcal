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

package cal_types

import (
	"errors"

	lib "github.com/ilius/libgostarcal"
)

// don't import "scal/utils"

type CalType interface {
	Name() string
	Desc() string
	Epoch() int
	MinMonthLen() uint8
	MaxMonthLen() uint8
	AvgYearLen() float64
	MonthNames() []string
	MonthNamesAb() []string
	IsLeap(year int) bool
	ToJd(date *lib.Date) int
	JdTo(jd int) *lib.Date
	GetMonthLen(year int, month uint8) uint8
}

type calTypeStruct struct {
	name string
	desc string

	isLeap      func(year int) bool
	toJd        func(date *lib.Date) int
	jdTo        func(jd int) *lib.Date
	getMonthLen func(year int, month uint8) uint8

	monthNames   []string
	monthNamesAb []string

	epoch       int
	minMonthLen uint8
	maxMonthLen uint8
	avgYearLen  float64
}

func (ct *calTypeStruct) Name() string {
	return ct.name
}

func (ct *calTypeStruct) Desc() string {
	return ct.desc
}

func (ct *calTypeStruct) Epoch() int {
	return ct.epoch
}

func (ct *calTypeStruct) MinMonthLen() uint8 {
	return ct.minMonthLen
}

func (ct *calTypeStruct) MaxMonthLen() uint8 {
	return ct.maxMonthLen
}

func (ct *calTypeStruct) AvgYearLen() float64 {
	return ct.avgYearLen
}

func (ct *calTypeStruct) MonthNames() []string {
	return ct.monthNames
}

func (ct *calTypeStruct) MonthNamesAb() []string {
	return ct.monthNamesAb
}

func (ct *calTypeStruct) IsLeap(year int) bool {
	return ct.isLeap(year)
}

func (ct *calTypeStruct) ToJd(date *lib.Date) int {
	return ct.toJd(date)
}

func (ct *calTypeStruct) JdTo(jd int) *lib.Date {
	return ct.jdTo(jd)
}

func (ct *calTypeStruct) GetMonthLen(year int, month uint8) uint8 {
	return ct.getMonthLen(year, month)
}

var CalTypesList []CalType
var CalTypesMap = make(map[string]CalType)

func RegisterCalType(
	name string,
	desc string,
	epoch int,
	minMonthLen uint8,
	maxMonthLen uint8,
	avgYearLen float64,
	monthNames []string,
	monthNamesAb []string,
	isLeap func(year int) bool,
	toJd func(date *lib.Date) int,
	jdTo func(jd int) *lib.Date,
	getMonthLen func(year int, month uint8) uint8,
) {
	calType := &calTypeStruct{
		name:         name,
		desc:         desc,
		epoch:        epoch,
		minMonthLen:  minMonthLen,
		maxMonthLen:  maxMonthLen,
		avgYearLen:   avgYearLen,
		monthNames:   monthNames,
		monthNamesAb: monthNamesAb,
		isLeap:       isLeap,
		toJd:         toJd,
		jdTo:         jdTo,
		getMonthLen:  getMonthLen,
	}
	CalTypesList = append(CalTypesList, calType)
	CalTypesMap[name] = calType
}

func invalidCalType(calTypeName string) error {
	return errors.New("invalid calendar type '" + calTypeName + "'")
}

func GetCalType(calTypeName string) (CalType, error) {
	calType, calTypeOk := CalTypesMap[calTypeName]
	if !calTypeOk {
		return nil, invalidCalType(calTypeName)
	}
	return calType, nil
}

func Convert(date *lib.Date, fromTypeName string, toTypeName string) (*lib.Date, error) {
	fromType, fromOk := CalTypesMap[fromTypeName]
	toType, toOk := CalTypesMap[toTypeName]
	if !fromOk {
		return nil, invalidCalType(fromTypeName)
	}
	if !toOk {
		return nil, invalidCalType(toTypeName)
	}
	return toType.JdTo(fromType.ToJd(date)), nil
}

func ToJd(date *lib.Date, calTypeName string) (int, error) {
	calType, calTypeOk := CalTypesMap[calTypeName]
	if !calTypeOk {
		return 0, invalidCalType(calTypeName)
	}
	return calType.ToJd(date), nil
}

func JdTo(jd int, calTypeName string) (*lib.Date, error) {
	calType, calTypeOk := CalTypesMap[calTypeName]
	if !calTypeOk {
		return nil, invalidCalType(calTypeName)
	}
	return calType.JdTo(jd), nil
}
