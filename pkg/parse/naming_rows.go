//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what namingRows — the rows of the closings given whose kinds take in the naming one (이름), in their order
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
)

func namingRows(rows []Marker) []Marker {
	var out []Marker
	for _, m := range rows {
		if m.Kind&rule.KindNoun != 0 {
			out = append(out, m)
		}
	}
	return out
}
