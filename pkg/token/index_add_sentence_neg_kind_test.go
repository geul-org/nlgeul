//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestIndexAddSentenceNegKind — the index of a reader holds the rule of rebuttal ruleNegKind twice after the rule of the -지 form before 않다, 못하다 or 말다 (linkBeforeNegRule), both with the force forceUngrammatical: the one that speaks of the adjective 못하 first, with the auxiliary verbs of the closed set among the verbs before it, then the one that speaks of the verb 못하
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestIndexAddSentenceNegKind(t *testing.T) {
	var ix index
	ix.addSentence()
	neg, kinds := -1, []int{}
	for i, s := range ix.sentence {
		if s.name == ruleLinkBeforeNeg {
			neg = i
		}
		if s.name == ruleNegKind && s.force == forceUngrammatical {
			kinds = append(kinds, i)
		}
	}
	if neg < 0 || len(kinds) != 2 || kinds[0] <= neg {
		t.Fatal(neg, kinds)
	}
	adj, verb := ix.sentence[kinds[0]], ix.sentence[kinds[1]]
	if !adj.about(rule.Near{Stem: "못하", Kind: "열쇠:형용사"}) || !verb.about(rule.Near{Stem: "못하", Kind: "열쇠:동사"}) {
		t.Fatal("the adjective's rule first, then the verb's")
	}
	aux := rule.Around{Self: rule.Near{Stem: "못하", Kind: "열쇠:형용사"}, Before: []rule.Near{{Stem: "있", Kind: "보조 용언:있", Marks: "맺음:지"}}}
	if !graphHolds(t, adj, aux) {
		t.Fatal("an auxiliary verb of the closed set is a verb before it")
	}
}
