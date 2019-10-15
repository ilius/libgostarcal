package rules_lib

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
)

const R_dayTimeRange = "dayTimeRange"

func init() {
	checker := func(value interface{}) bool {
		v, ok := value.(*lib.HMSRange)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_dayTimeRange,
				value,
				value,
			))
		}
		return v.IsValid()
	}
	RegisterRuleType(
		6,
		R_dayTimeRange,
		T_HMSRange,
		&checker,
	)
}
