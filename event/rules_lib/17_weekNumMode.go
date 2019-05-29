package rules_lib

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
)

const R_weekNumMode = "weekNumMode"

func init() {
	checker := func(value interface{}) bool {
		v, ok := value.(string)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_weekNumMode,
				value,
				value,
			))
		}
		switch v {
		case lib.ODD:
			return true
		case lib.EVEN:
			return true
		case lib.ANY:
			return true
		}
		return false
	}
	RegisterRuleType(
		17,
		R_weekNumMode,
		T_string,
		&checker,
	)
}
