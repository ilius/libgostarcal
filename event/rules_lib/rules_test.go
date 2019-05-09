package rules_lib

import (
	"github.com/ilius/is"
	"testing"
)

type RuleTestCase struct {
	Type     string
	Value    string
	DecodeOk bool
	CheckOk  bool
}

func TestRules(t *testing.T) {
	test := func(t *testing.T, typ string, value string, decodeOk bool, checkOk bool) {
		is := is.New(t).MsgSep(", ").Msg("type=%v, value=%#v, decodeOk=%v, checkOk=%v", typ, value, decodeOk, checkOk)
		t.Helper()
		model := EventRuleModel{
			Type:  typ,
			Value: value,
		}
		rule, err := model.Decode()
		is.PrependMsg("mismatch decodeOk").Equal(err == nil, decodeOk)
		if err != nil {
			return
		}
		is = is.PrependMsg("rule.Value=%#v", rule.Value)
		checkOkActual := rule.Check()
		is.PrependMsg("mismatch checkOk").Equal(checkOkActual, checkOk)
	}
	test(t, "cycleDays", "10", true, false)
	test(t, "cycleDays", "-1", true, false)
	test(t, "cycleDays", "2f", false, false)
	test(t, "cycleLen", "90 23:55:55", true, true)
	test(t, "cycleLen", "1 23:55:55", true, true)
	test(t, "cycleLen", "10 -1:55:55", true, false)
	test(t, "cycleLen", "10", false, false)
	test(t, "cycleLen", "10 a:b:c", false, false)
	test(t, "date", "2000/1/1", true, true)
	test(t, "date", "-1000/1/1", true, true)
	test(t, "date", "-1000/12/31", true, true)
	test(t, "date", "2000/13/31", true, false)
	test(t, "date", "2000/12/-1", true, false)
	test(t, "date", "2000-1-1", false, false)
	test(t, "date", "2000", false, false)
	test(t, "ex_dates", "2000/1/1 2010/12/1 -1000/1/31", true, true)
	test(t, "ex_dates", "aa/1/1 2010/12/1 -1000/1/31", false, false)
	test(t, "day", "1", true, true)
	test(t, "day", "1 5 10 15 30", true, true)
	test(t, "day", "31", true, true)
	test(t, "day", "0", true, false)
	test(t, "day", "50", true, false)
	test(t, "day", "1 50 5 10 15 30", true, false)
	test(t, "day", "ff", false, false)
	test(t, "ex_day", "1", true, true)
	test(t, "ex_day", "1 5 10 15 30", true, true)
	test(t, "ex_day", "31", true, true)
	test(t, "ex_day", "0", true, false)
	test(t, "ex_day", "50", true, false)
	test(t, "ex_day", "1 50 5 10 15 30", true, false)
	test(t, "ex_day", "ff", false, false)
	test(t, "dayTime", "22:30:55", true, true)
	test(t, "dayTime", "24:30:55", true, false)
	test(t, "dayTime", "20:60:55", true, false)
	test(t, "dayTime", "20:55:61", true, false)
	test(t, "dayTime", "20:55:-1", true, false)
	test(t, "dayTime", "205501", false, false)
	test(t, "dayTime", "ab", false, false)
	test(t, "dayTimeRange", "22:30:55 23:30:55", true, true)
	test(t, "dayTimeRange", "22:30:55 24:30:55", true, false)
	test(t, "dayTimeRange", "22:30:55-23:30:55", false, false)
	test(t, "duration", "3.1 d", true, true)
	test(t, "duration", "-1 d", true, false)
	test(t, "duration", "-1.5 w", true, false)
	test(t, "duration", "1d", false, false)
	test(t, "end", "2016/12/31 23:55:55", true, true)
	test(t, "end", "2016/12/31 23:55:61", true, false)
	test(t, "end", "2016/12/31-23:55:55", false, false)
	test(t, "month", "1", true, true)
	test(t, "month", "12", true, true)
	test(t, "month", "0", true, false)
	test(t, "month", "bb", false, false)
	test(t, "ex_month", "1", true, true)
	test(t, "ex_month", "12", true, true)
	test(t, "ex_month", "0", true, false)
	test(t, "ex_month", "bb", false, false)
	test(t, "start", "-1/11/30 22:50:00", true, true)
	test(t, "start", "-1/0/30 22:50:00", true, false)
	test(t, "start", "-1/11/30", false, false)
	test(t, "weekDay", "0 2 4 6", true, true)
	test(t, "weekDay", "-1 2 4 6", true, false)
	test(t, "weekDay", "0 2 4 7", true, false)
	test(t, "weekDay", "a 2 4 7", false, false)
	test(t, "weekNumMode", "odd", true, true)
	test(t, "weekNumMode", "even", true, true)
	test(t, "weekNumMode", "any", true, true)
	test(t, "weekNumMode", "foo", true, false)
	test(t, "year", "1000 100 0 1", true, true)
	test(t, "year", "-(600-500) -400 -300", true, true)
	test(t, "year", "-600--500 -400 -300", true, true)
	test(t, "year", "ff 1000 100 0 1", false, false)
	test(t, "ex_year", "1000 100 0 1", true, true)
	test(t, "ex_year", "ff 1000 100 0 1", false, false)
	test(t,
		"weekMonth",
		"{\"weekIndex\": 4, \"weekDay\": 6, \"month\": 12}",
		true,
		true,
	)
	test(t,
		"weekMonth",
		"{\"weekIndex\": 1, \"weekDay\": 1, \"month\": 0}",
		true,
		true,
	)
	test(t,
		"weekMonth",
		"{\"weekIndex\": 5, \"weekDay\": 6, \"month\": 12}",
		true,
		false,
	)
	test(t,
		"weekMonth",
		"{\"weekIndex\": 0, \"weekDay\": 7, \"month\": 12}",
		true,
		false,
	)
	test(t,
		"weekMonth",
		"\"weekIndex\": 0, \"weekDay\": 0, \"month\": 0",
		false,
		false,
	)
}
