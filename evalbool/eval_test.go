package evalbool

import "testing"

type eData struct {
	in  string
	out bool
}

func TestEvalBool(t *testing.T) {
	data := []eData{
		eData{
			in:  "T!",
			out: false,
		},
		eData{
			in:  "TT&",
			out: true,
		},
		eData{
			in:  "TF&",
			out: false,
		},
		eData{
			in:  "TTF|&",
			out: true,
		},
		eData{
			in:  "tf|tf|&",
			out: true,
		},
		eData{
			in:  "ff&",
			out: false,
		},
		eData{
			in:  "ff^",
			out: false,
		},
		eData{
			in:  "TT^",
			out: false,
		},
		eData{
			in:  "TF^",
			out: true,
		},
		eData{
			in:  "FT^",
			out: true,
		},
	}

	for _, v := range data {
		e, err := EvalRPN([]byte(v.in))
		if err != nil {
			t.Fatal(err)
		}
		if e != v.out {
			t.Fatalf("[%s] expected %t, got %t", v.in, v.out, e)
		}
	}
}

func TestEval(t *testing.T) {
	data := []eData{
		eData{
			in:  "t&t",
			out: true,
		},
		eData{
			in:  "t&f",
			out: false,
		},
		eData{
			in:  "(t|f)&(f|f)",
			out: false,
		},
	}
	for _, v := range data {
		e, err := Eval([]byte(v.in))
		if err != nil {
			t.Fatal(err)
		}
		if e != v.out {
			t.Fatalf("[%s] expected %t, got %t", v.in, v.out, e)
		}
	}
}
