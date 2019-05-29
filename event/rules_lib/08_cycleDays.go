package rules_lib

import "fmt"

const R_cycleDays = "cycleDays"

func init() {
	checker := func(value interface{}) bool {
		v, ok := value.(int)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_cycleDays,
				value,
				value,
			))
		}
		return v > 0
	}
	RegisterRuleType(
		8,
		R_cycleDays,
		T_int,
		&checker,
	)
}
