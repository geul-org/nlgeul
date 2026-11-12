//ff:func feature=hypo type=helper control=iteration dimension=1
//ff:what Choices — the bundles of a sentence a hypothesis picks a reading from: those nlgeul token prints, in order (the bundles of readings asleep are left out)
package hypo

import "github.com/geul-org/nlgeul/internal/token"

func Choices(bundles []token.Bundle) []token.Bundle {
	var out []token.Bundle
	for _, b := range bundles {
		if !b.Asleep() {
			out = append(out, b)
		}
	}
	return out
}
