//ff:func feature=hypo type=helper control=iteration dimension=1
//ff:what Picked — the readings of the hypothesis with the given number, one for every bundle in order
package hypo

import "github.com/geul-org/nlgeul/internal/token"

// The number is read as digits, one for each bundle, the last bundle's
// digit lowest: each digit is which reading of that bundle is picked.
func Picked(choices []token.Bundle, number int) []token.Branch {
	out := make([]token.Branch, len(choices))
	for i := len(choices) - 1; i >= 0; i-- {
		n := len(choices[i].Branches)
		out[i], number = choices[i].Branches[number%n], number/n
	}
	return out
}
