//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestInterjectionAloneRuleGraph — the graph of interjectionAloneRule evaluated by hand on rule.Around: the interjection 흠 written apart stands at the head of a sentence, before a comma and a noun, and with a sign that closes written onto it; it falls right after a modifying form, a nominal with the adnominal case particle and a determiner of the closed set, when it is written onto a word, and when a noun or an opening sign is written onto it (흠(…), a gloss of a name)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestInterjectionAloneRuleGraph(t *testing.T) {
	apart := rule.Near{Stem: "흠", Kind: "열쇠:홀로"}
	comma := rule.Near{Stem: ",", Kind: "쉼표", Glued: true}
	noun := rule.Near{Stem: "비용", Kind: "열쇠:체언", Last: "주격:이", Marks: "주격:이", Modifiable: true}
	r := interjectionAloneRule(map[string]word{"흠": {kinds: kindNominal | kindAlone}}, map[string]int{})
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before a comma", rule.Around{Self: apart, After: []rule.Near{comma}}, true},
		{"before a noun", rule.Around{Self: apart, After: []rule.Near{noun}}, true},
		{"a closing sign written onto it", rule.Around{Self: apart, After: []rule.Near{{Stem: "”", Kind: "닫기:1", Glued: true}}}, true},
		{"after a modifying form", rule.Around{Self: apart, Before: []rule.Near{{Stem: "좋", Kind: "열쇠:형용사", Last: "꾸밈:은", Pred: true}}}, false},
		{"after the adnominal case", rule.Around{Self: apart, Before: []rule.Near{{Stem: "학교", Kind: "열쇠:체언", Last: "관형격:의"}}}, false},
		{"after a determiner", rule.Around{Self: apart, Before: []rule.Near{{Stem: "그", Kind: "관형사:그"}}}, false},
		{"onto a word", rule.Around{Self: rule.Near{Stem: "흠", Kind: "열쇠:홀로", Glued: true}, Before: []rule.Near{noun}, After: []rule.Near{comma}}, false},
		{"a noun written onto it", rule.Around{Self: apart, After: []rule.Near{{Stem: "비용", Kind: "열쇠:체언", Modifiable: true, Glued: true}}}, false},
		{"an opening sign written onto it", rule.Around{Self: apart, After: []rule.Near{{Stem: "(", Kind: "열기:1", Glued: true}}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
