//ff:func feature=token type=reader control=iteration dimension=1
//ff:what signFacts — the facts of the sign after the k-th run of a chain: those of a dot, taken over the runs that the same sign joins around it, and those runs
package token

import "github.com/geul-org/koparse/pkg/rule"

// signs[i] is the sign after the i-th run.
func signFacts(labels []string, signs []rune, k int) rule.Mark {
	lo, hi := k, k+1
	for lo > 0 && signs[lo-1] == signs[k] {
		lo--
	}
	for hi < len(signs) && signs[hi] == signs[k] {
		hi++
	}
	m := dotFacts(labels[lo:hi+1], k-lo)
	m.Parts = labels[lo : hi+1]
	return m
}
