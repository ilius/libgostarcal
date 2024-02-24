package rules_lib

import (
	"fmt"

	"github.com/ilius/libgostarcal/utils"
)

const R_weekDay = "weekDay"

func init() {
	checker := func(value any) bool {
		list, ok := value.([]int)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_weekDay,
				value,
				value,
			))
		}
		return utils.WeekDayListIsValid(list)
	}
	RegisterRuleType(
		10,
		R_weekDay,
		T_int_list,
		&checker,
	)
}
