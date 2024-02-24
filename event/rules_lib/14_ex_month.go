package rules_lib

import (
	"fmt"

	"github.com/ilius/libgostarcal/utils"
)

const R_ex_month = "ex_month"

func init() {
	checker := func(value any) bool {
		list, ok := value.([]int)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_ex_month,
				value,
				value,
			))
		}
		return utils.MonthListIsValid(list)
	}
	RegisterRuleType(
		14,
		R_ex_month,
		T_int_range_list,
		&checker,
	)
}
