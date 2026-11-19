//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestIndexAddSentenceAdnHead — the index of a reader holds the rule of rebuttal ruleAdnHead twice after the rule of 이라고 after 것 (thingQuoteRule), both with the force forceUngrammatical: the one judged on the adnominal first (it waits for the words after it), then the one judged on the word after it
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestIndexAddSentenceAdnHead(t *testing.T) {
	var ix index
	ix.addSentence()
	quote, heads := -1, []int{}
	for i, s := range ix.sentence {
		if s.name == ruleThingQuote {
			quote = i
		}
		if s.name == ruleAdnHead && s.force == forceUngrammatical {
			heads = append(heads, i)
		}
	}
	if quote < 0 || len(heads) != 2 || heads[0] <= quote {
		t.Fatal(quote, heads)
	}
	adn, after := ix.sentence[heads[0]], ix.sentence[heads[1]]
	if !adn.about(rule.Near{Stem: "이", Kind: "관형사:이"}) || adn.waits == nil || !after.about(rule.Near{Stem: "채", Kind: "열쇠:부사"}) || after.waits != nil {
		t.Fatal("the adnominal's rule first, then the one of the word after it")
	}
}
