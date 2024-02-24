package rules_lib

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
)

const R_cycleLen = "cycleLen"

func init() {
	checker := func(value any) bool {
		v, ok := value.(*lib.DHMS)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_cycleLen,
				value,
				value,
			))
		}
		return v.IsValid()
	}
	RegisterRuleType(
		7,
		R_cycleLen,
		T_DHMS,
		&checker,
	)
}
