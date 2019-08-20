package libscal

import (
	"testing"

	"github.com/ilius/is"
)

func TestHMS_FloatHour(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	hms := HMS{12, 59, 5}
	is = is.AddMsg("hms=%v", hms)
	fh := hms.GetFloatHour()
	is.Equal(12.98472222222222, fh)
	hms2 := FloatHourToHMS(fh)
	is.Equal("12:59:05", hms2.String())
}

func TestParseHMS(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	str := "12:01:01"
	is = is.AddMsg("str=%v", str)
	obj, err := ParseHMS(str)
	is.NotErr(err)
	is.Equal(str, obj.String())
}

func TestParseDHMS(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	str := "90 12:01:01"
	is = is.AddMsg("str=%v", str)
	obj, err := ParseDHMS(str)
	is.NotErr(err)
	is.Equal(str, obj.String())
}
