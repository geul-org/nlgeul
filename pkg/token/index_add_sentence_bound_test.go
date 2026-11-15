//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestIndexAddSentenceBound — the index of a reader holds the rules of dependent nouns and of the nouns after a modifier after quotedRule, in this order: boundBeforeRule, nounUnmodifiedRule, boundAfterRule, nounListedRule, nounHeadedRule, nounModifiedRule (other rules may be between them)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceBound(t *testing.T) {
	r := streamReader(t, "words-bound.tsv")
	var names []string
	for _, s := range r.ix.sentence {
		names = append(names, s.name)
	}
	at := slices.Index(names, ruleQuoted)
	for _, name := range []string{ruleBoundBefore, ruleNounUnmodified, ruleBoundAfter, ruleNounListed, ruleNounHeaded, ruleNounModified} {
		next := slices.Index(names, name)
		if next <= at {
			t.Fatalf("%s is missing or out of order in %v", name, names)
		}
		at = next
	}
}
