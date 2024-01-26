package registry

import (
	"github.com/ilius/libgostarcal/cal_types"
	"github.com/ilius/libgostarcal/cal_types/ethiopian"
	"github.com/ilius/libgostarcal/cal_types/gregorian"
	"github.com/ilius/libgostarcal/cal_types/gregorian_proleptic"
	"github.com/ilius/libgostarcal/cal_types/hijri"
	"github.com/ilius/libgostarcal/cal_types/indian_national"
	"github.com/ilius/libgostarcal/cal_types/jalali"
	"github.com/ilius/libgostarcal/cal_types/julian"
)

func init() {
	cal_types.RegisterCalType(ethiopian.New())
	cal_types.RegisterCalType(gregorian.New())
	cal_types.RegisterCalType(gregorian_proleptic.New())
	cal_types.RegisterCalType(hijri.New())
	cal_types.RegisterCalType(indian_national.New())
	cal_types.RegisterCalType(jalali.New())
	cal_types.RegisterCalType(julian.New())
}
