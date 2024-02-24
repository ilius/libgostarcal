package rules_lib

import (
	"fmt"

	"github.com/ilius/libgostarcal/utils"
)

const R_day = "day"

func init() {
	checker := func(value any) bool {
		list, ok := value.([]int)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_day,
				value,
				value,
			))
		}
		return utils.DayListIsValid(list)
	}
	RegisterRuleType(
		15,
		R_day,
		T_int_range_list,
		&checker,
	)
}
