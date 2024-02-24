package rules_lib

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
)

const R_dayTime = "dayTime"

func init() {
	checker := func(value any) bool {
		v, ok := value.(*lib.HMS)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_dayTime,
				value,
				value,
			))
		}
		return v.IsValid()
	}
	RegisterRuleType(
		5,
		R_dayTime,
		T_HMS,
		&checker,
	)
}
