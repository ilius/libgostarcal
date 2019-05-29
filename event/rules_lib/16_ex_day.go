package rules_lib

import (
	"fmt"

	"github.com/ilius/libgostarcal/utils"
)

const R_ex_day = "ex_day"

func init() {
	checker := func(value interface{}) bool {
		list, ok := value.([]int)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_ex_day,
				value,
				value,
			))
		}
		return utils.DayListIsValid(list)
	}
	RegisterRuleType(
		16,
		R_ex_day,
		T_int_range_list,
		&checker,
	)
}
