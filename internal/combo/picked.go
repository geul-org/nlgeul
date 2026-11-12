//ff:func feature=combo type=helper control=iteration dimension=1
//ff:what Picked — the candidates of the combination with the given number, one for every contest in order
package combo

import "github.com/geul-org/nlgeul/internal/token"

// The number is read as digits, one for each contest, the last contest's
// digit lowest: each digit is which candidate of that contest is picked.
func Picked(choices []token.Contest, number int) []token.Branch {
	out := make([]token.Branch, len(choices))
	for i := len(choices) - 1; i >= 0; i-- {
		n := len(choices[i].Branches)
		out[i], number = choices[i].Branches[number%n], number/n
	}
	return out
}
