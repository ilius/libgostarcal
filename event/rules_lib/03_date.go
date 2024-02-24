package rules_lib

import (
	"fmt"

	lib "github.com/ilius/libgostarcal"
)

const R_date = "date"

func init() {
	checker := func(value any) bool {
		v, ok := value.(*lib.Date)
		if !ok {
			panic(fmt.Errorf(
				"%s rule value checker: type conversion failed, value=%v with type %T\n",
				R_date,
				value,
				value,
			))
		}
		return v.IsValid()
	}
	RegisterRuleType(
		3,
		R_date,
		T_Date,
		&checker,
	)
}
