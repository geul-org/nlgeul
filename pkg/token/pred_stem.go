//ff:func feature=token type=helper control=sequence
//ff:what predStem — does the candidate read its stem as a predicate's: a verb or an adjective of the word table, or an auxiliary verb of the closed set (not a nominal, a new word, nor the copula's piece after a nominal)
package token

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func predStem(n rule.Near) bool {
	return slices.Contains(predStemKinds, n.Kind) || strings.HasPrefix(n.Kind, auxSet+":")
}
