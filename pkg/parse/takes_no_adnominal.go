//ff:func feature=parse type=helper control=selection
//ff:what takesNoAdnominal — is a word's candidate one no adnominal modifies: a predicate's stem (the word table's verb or adjective, an auxiliary verb of the closed set) none of whose markers is a modifying or a naming closing, an adverb (the word table's or a conjunctive one of the closed set) or an interjection
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// School grammar: an adnominal modifies the nominal after it. A
// predicate that ends otherwise than in a modifying form (which may
// stand in a row of adnominals) or a naming form (which stands as a
// nominal), an adverb and an interjection are no nominal.
func takesNoAdnominal(n rule.Near) bool {
	switch {
	case n.Kind == keyKind+":"+kindLabel[kindAdverb], n.Kind == keyKind+":"+kindLabel[kindAlone], strings.HasPrefix(n.Kind, conjAdverbLabel+":"):
		return true
	case n.Kind != keyKind+":"+kindLabel[kindVerb] && n.Kind != keyKind+":"+kindLabel[kindAdj] && !strings.HasPrefix(n.Kind, auxSet+":"):
		return false
	}
	return !strings.Contains(n.Marks, adnPrefix) && !strings.Contains(n.Marks, namingPrefix)
}
