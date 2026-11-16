//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAdverbAloneRuleGraph — the graph of adverbAloneRule evaluated by hand on rule.Around: the adverb 모두 written apart stands at the head of a sentence, before a noun, before a predicate, before a sign and at the end, and with a sign or a predicate written onto it; it falls right after a modifying form, a nominal with the adnominal case particle and a determiner of the closed set, when it is written onto a word and when a nominal is written onto it; written onto an opening sign it stands
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdverbAloneRuleGraph(t *testing.T) {
	apart := rule.Near{Stem: "모두", Kind: "열쇠:부사"}
	glued := rule.Near{Stem: "모두", Kind: "열쇠:부사", Glued: true}
	noun := rule.Near{Stem: "학교", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에", Modifiable: true}
	verb := rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:다", Pred: true}
	open := rule.Near{Stem: "“", Kind: "열기:1"}
	r := adverbAloneRule(map[string]word{"모두": {kinds: kindNominal | kindAdverb}})
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"at the head before a noun", rule.Around{Self: apart, After: []rule.Near{noun}}, true},
		{"before a predicate", rule.Around{Self: apart, Before: []rule.Near{{Stem: "학생", Kind: "열쇠:체언", Last: "주격:이"}}, After: []rule.Near{verb}}, true},
		{"before a comma", rule.Around{Self: apart, After: []rule.Near{{Stem: ",", Kind: "쉼표", Glued: true}}}, true},
		{"at the end", rule.Around{Self: apart, Before: []rule.Near{verb}}, true},
		{"a predicate written onto it", rule.Around{Self: apart, After: []rule.Near{{Stem: "가", Kind: "열쇠:동사", Last: "맺음:다", Pred: true, Glued: true}}}, true},
		{"after a modifying form", rule.Around{Self: apart, Before: []rule.Near{{Stem: "좋", Kind: "열쇠:형용사", Last: "꾸밈:은", Pred: true}}}, false},
		{"after the adnominal case", rule.Around{Self: apart, Before: []rule.Near{{Stem: "학교", Kind: "열쇠:체언", Last: "관형격:의"}}}, false},
		{"after a determiner", rule.Around{Self: apart, Before: []rule.Near{{Stem: "이", Kind: "관형사:이"}}}, false},
		{"onto a word", rule.Around{Self: glued, Before: []rule.Near{noun}, After: []rule.Near{verb}}, false},
		{"onto an opening sign", rule.Around{Self: glued, Before: []rule.Near{open}, After: []rule.Near{verb}}, true},
		{"a nominal written onto it", rule.Around{Self: apart, After: []rule.Near{{Stem: "출자", Kind: "열쇠:체언", Modifiable: true, Glued: true}}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
