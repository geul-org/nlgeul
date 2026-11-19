//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestIndexAddSentenceOpenMarker — the index of a reader holds the rule of rebuttal ruleOpenMarker once, after the rules of an adnominal and the word after it (ruleAdnHead), with the force forceUngrammatical, speaking of markers alone
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestIndexAddSentenceOpenMarker(t *testing.T) {
	var ix index
	ix.addSentence()
	head, opens := -1, []int{}
	for i, s := range ix.sentence {
		if s.name == ruleAdnHead {
			head = i
		}
		if s.name == ruleOpenMarker && s.force == forceUngrammatical {
			opens = append(opens, i)
		}
	}
	if head < 0 || len(opens) != 1 || opens[0] <= head {
		t.Fatal(head, opens)
	}
	if !ix.sentence[opens[0]].about(rule.Near{Last: "목적격:을", Marks: "목적격:을"}) {
		t.Fatal("markers alone are asked")
	}
}
