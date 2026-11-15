//ff:func feature=token type=helper control=iteration dimension=1
//ff:what nearsOn — the words after a seat in the line, the nearest first
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func nearsOn(seats []seat) []rule.Near {
	out := make([]rule.Near, 0, len(seats))
	for _, st := range seats {
		out = append(out, st.near)
	}
	return out
}
