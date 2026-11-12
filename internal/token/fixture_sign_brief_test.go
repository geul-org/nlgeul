//ff:func feature=token type=helper control=iteration dimension=1
//ff:what signBrief — a bundle of signs in short: each branch as the first part's names, whether it is a stem, its depth and where its word ends, the branches joined by "; "
package token

import (
	"fmt"
	"strings"
)

func signBrief(b Bundle) string {
	var out []string
	for _, br := range b.Branches {
		p := br.Words[0].Parts[0]
		out = append(out, fmt.Sprintf("%q/stem=%v/depth=%d/end=%d", namesText(p.Names), p.Stem, p.Depth, br.Words[0].End))
	}
	return strings.Join(out, "; ")
}
