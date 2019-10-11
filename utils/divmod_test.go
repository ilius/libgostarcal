package utils

import (
	"testing"
	"github.com/ilius/is"
)

func TestDivmod(t *testing.T) {
	is := is.New(t).MsgSep(", ")
	testMap := map[[2]int][2]int{
		{-20, -10}: {2, 0},
		{20, -10}:  {-2, 0},
		{-20, 10}:  {-2, 0},
		{-23, -10}: {2, -3},
		{-23, 10}:  {-3, 7},
		{-13, 10}:  {-2, 7},
		{12, 10}:   {1, 2},
		{22, 10}:   {2, 2},
	}
	for ab, dm := range testMap {
		a := ab[0]
		b := ab[1]
		divExpect := dm[0]
		modExpect := dm[1]

		is.AddMsg("Div: a=%v, b=%v", a, b).Equal(Div(a, b), divExpect)
		is.AddMsg("Mod: a=%v, b=%v", a, b).Equal(Mod(a, b), modExpect)

		div, mod := Divmod(a, b)
		is.AddMsg("Divmod: mismatch div, a=%v, b=%v", a, b).Equal(div, divExpect)
		is.AddMsg("Divmod: mismatch mod, a=%v, b=%v", a, b).Equal(mod, modExpect)
	}
}

