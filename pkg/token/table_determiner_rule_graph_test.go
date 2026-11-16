//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestTableDeterminerRuleGraph — the graph of tableDeterminerRule evaluated by hand on rule.Around: the determiner 총 written apart stands before a noun, a number, a determiner and a modifying form, written onto an opening sign too, and with a noun written onto it; it falls before an adverb, a predicate, a sign and at the end of the sentence, and written onto a word
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestTableDeterminerRuleGraph(t *testing.T) {
	apart := rule.Near{Stem: "총", Kind: "열쇠:관형사"}
	glued := rule.Near{Stem: "총", Kind: "열쇠:관형사", Glued: true}
	noun := rule.Near{Stem: "비용", Kind: "열쇠:체언", Last: "주격:이", Marks: "주격:이", Modifiable: true}
	r := tableDeterminerRule(map[string]word{"총": {kinds: kindNominal | kindDeterminer}}, map[string]int{})
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before a noun", rule.Around{Self: apart, After: []rule.Near{noun}}, true},
		{"before a number", rule.Around{Self: apart, After: []rule.Near{{Stem: "3", Kind: "수", Last: "단위:명", Modifiable: true}}}, true},
		{"before a determiner", rule.Around{Self: apart, After: []rule.Near{{Stem: "한", Kind: "관형사:한", Modifiable: true}}}, true},
		{"before a modifying form", rule.Around{Self: apart, After: []rule.Near{{Stem: "좋", Kind: "열쇠:형용사", Last: "꾸밈:은", Pred: true, Modifiable: true}}}, true},
		{"a noun written onto it", rule.Around{Self: apart, After: []rule.Near{{Stem: "비용", Kind: "열쇠:체언", Modifiable: true, Glued: true}}}, true},
		{"onto an opening sign", rule.Around{Self: glued, Before: []rule.Near{{Stem: "(", Kind: "열기:1"}}, After: []rule.Near{noun}}, true},
		{"before an adverb", rule.Around{Self: apart, After: []rule.Near{{Stem: "다시", Kind: "열쇠:부사"}, noun}}, false},
		{"before a predicate", rule.Around{Self: apart, After: []rule.Near{{Stem: "좋", Kind: "열쇠:형용사", Last: "맺음:다", Pred: true}}}, false},
		{"before a sign", rule.Around{Self: apart, After: []rule.Near{{Stem: ",", Kind: "쉼표", Glued: true}}}, false},
		{"at the end", rule.Around{Self: apart}, false},
		{"onto a word", rule.Around{Self: glued, Before: []rule.Near{{Stem: "학교", Kind: "열쇠:체언", Last: "부사격:에"}}, After: []rule.Near{noun}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
