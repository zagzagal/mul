package evalbool

import (
	"container/list"
)

// Shunt takes the byte array form of a boolian expression and puts it into
// reverse polish notation
func Shunt(b []byte) []byte {
	output := list.New()
	opp := list.New()

	for _, v := range b {
		switch v {
		case '&', '|', '^':
			e := opp.Back()
			for opp.Len() > 0 && e.Value.(byte) != '(' {
				//output.Push(opp.Pop())
				opp.Remove(e)
				output.PushBack(e.Value.(byte))
				e = opp.Back()
			}
			opp.PushBack(v)
		case '(', '!':
			opp.PushBack(v)
		case ')':
			e := opp.Back()
			for opp.Len() > 0 && e.Value.(byte) != '(' {
				opp.Remove(e)
				output.PushBack(e.Value.(byte))

				e = opp.Back()
			}
			if e.Value.(byte) == '(' {
				e := opp.Back()
				opp.Remove(e)
			}
		case ' ':
		default:
			output.PushBack(v)
		}
	}
	for opp.Len() > 0 {
		e := opp.Back()
		opp.Remove(e)
		output.PushBack(e.Value.(byte))
	}
	return listBytes(output)
}

func listBytes(l *list.List) []byte {
	out := make([]byte, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		out[i] = e.Value.(byte)
		i++
	}
	return out
}
