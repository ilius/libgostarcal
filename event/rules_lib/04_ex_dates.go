package rules_lib

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
)

const R_ex_dates = "ex_dates"

func init() {
	checker := func(value any) bool {
		list, ok := value.([]*lib.Date)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_ex_dates,
				value,
				value,
			))
		}
		for _, date := range list {
			if !date.IsValid() {
				return false
			}
		}
		return true
	}
	RegisterRuleType(
		4,
		R_ex_dates,
		T_Date_list,
		&checker,
	)
}
