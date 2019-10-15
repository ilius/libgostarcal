package rules_lib

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
)

const R_end = "end"

func init() {
	checker := func(value interface{}) bool {
		v, ok := value.(*lib.DateHMS)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_end,
				value,
				value,
			))
		}
		return v.IsValid()
	}
	RegisterRuleType(
		1,
		R_end,
		T_DateHMS,
		&checker,
	)
}
