package rules_lib

import "fmt"

const R_cycleWeeks = "cycleWeeks"

func init() {
	checker := func(value any) bool {
		v, ok := value.(int)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_cycleWeeks,
				value,
				value,
			))
		}
		return v > 0
	}
	RegisterRuleType(
		9,
		R_cycleWeeks,
		T_int,
		&checker,
	)
}
