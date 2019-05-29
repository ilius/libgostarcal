package rules_lib

import (
	"fmt"

	"github.com/ilius/libgostarcal/utils"
)

const R_duration = "duration"

func init() {
	checker := func(value interface{}) bool {
		v, ok := value.(utils.Duration)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_duration,
				value,
				value,
			))
		}
		return v.IsValid()
	}
	RegisterRuleType(
		2,
		R_duration,
		T_Duration,
		&checker,
	)
}
