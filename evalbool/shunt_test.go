package evalbool

import (
	"testing"
)

type tData struct {
	in  string
	out string
}

func TestShunt(t *testing.T) {
	data := []tData{
		tData{
			in:  "a&!b",
			out: "ab!&",
		},
		tData{
			in:  "!(a|b)",
			out: "ab|!",
		},
		tData{
			in:  "!a",
			out: "a!",
		},
		tData{
			in:  "a&b",
			out: "ab&",
		},
		tData{
			in:  "a|b",
			out: "ab|",
		},
		tData{
			in:  "a|b&c",
			out: "ab|c&",
		},
		tData{
			in:  "a | b",
			out: "ab|",
		},
		tData{
			in:  "(a|b)&(c|d)",
			out: "ab|cd|&",
		},
	}

	for _, v := range data {
		in := []byte(v.in)
		s := Shunt(in)
		if string(s) != v.out {
			t.Fatalf("%s != %s", s, v.out)
		}
	}
}
