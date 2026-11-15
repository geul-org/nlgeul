//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestIndexAddSentenceCopula — the rules of a sentence that read the copula and the particle of quotation (ruleCopulaEnd, ruleNumberQuote, ruleCopulaComma) are in the index after ruleQuoted, in this order
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceCopula(t *testing.T) {
	var ix index
	ix.addSentence()
	var names []string
	for _, s := range ix.sentence {
		names = append(names, s.name)
	}
	at := func(name string) int { return slices.Index(names, name) }
	if at(ruleQuoted) < 0 || at(ruleQuoted) >= at(ruleCopulaEnd) || at(ruleCopulaEnd) >= at(ruleNumberQuote) || at(ruleNumberQuote) >= at(ruleCopulaComma) {
		t.Fatal(names)
	}
}
