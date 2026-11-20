//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what nearsBack — the words before a seat, the nearest first: the seats before it from the last back, then the words judged earlier (behind, already the nearest first)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func nearsBack(seats []seat, behind []rule.Near) []rule.Near {
	out := make([]rule.Near, 0, len(seats)+len(behind))
	for i := len(seats) - 1; i >= 0; i-- {
		out = append(out, seats[i].near)
	}
	return append(out, behind...)
}
