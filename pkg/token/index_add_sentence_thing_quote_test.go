//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentenceThingQuote — the index of a reader holds the rule of rebuttal ruleThingQuote, with the force forceUngrammatical, after the rules of 못하다 after -지 (negKindRule)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceThingQuote(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	kind, quote := at(ruleNegKind), at(ruleThingQuote)
	if kind < 0 || quote <= kind || ix.sentence[quote].force != forceUngrammatical {
		t.Fatal(kind, quote, len(ix.sentence))
	}
}
