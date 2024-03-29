package rules_lib

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
)

const R_start = "start"

func init() {
	checker := func(value any) bool {
		v, ok := value.(*lib.DateHMS)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_start,
				value,
				value,
			))
		}
		return v.IsValid()
	}
	RegisterRuleType(
		0,
		R_start,
		T_DateHMS,
		&checker,
	)
}
