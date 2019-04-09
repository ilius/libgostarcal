package rules_lib

const R_year = "year"

// FIXME: how can we add support for negetive year numbers?
// without dropping support for year range

func init() {
	RegisterRuleType(
		11,
		R_year,
		T_int_range_list,
		nil,
	)
}
