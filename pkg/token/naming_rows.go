//ff:func feature=token type=reader control=iteration dimension=1
//ff:what namingRows — the rows of the closings given whose kinds take in the naming one (이름), in their order
package token

import (
	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func namingRows(rows []parse.Marker) []parse.Marker {
	var out []parse.Marker
	for _, m := range rows {
		if m.Kind&rule.KindNoun != 0 {
			out = append(out, m)
		}
	}
	return out
}
