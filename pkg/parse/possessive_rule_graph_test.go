//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPossessiveRuleGraph — the graph of possessiveRule evaluated by hand on rule.Around: 제 written apart stands before a noun, a modifying form and an opening sign, and written onto an opening sign; it falls before a unit, an adverb, a predicate, a sign and at the end of the sentence, and written onto a word
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPossessiveRuleGraph(t *testing.T) {
	apart := rule.Near{Stem: "제", Kind: "열쇠:체언"}
	glued := rule.Near{Stem: "제", Kind: "열쇠:체언", Glued: true}
	noun := rule.Near{Stem: "집", Kind: "열쇠:체언", Last: "부사격:으로", Marks: "부사격:으로", Modifiable: true}
	r := possessiveRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before a noun", rule.Around{Self: apart, After: []rule.Near{noun}}, true},
		{"before a modifying form", rule.Around{Self: apart, After: []rule.Near{{Stem: "좋", Kind: "열쇠:형용사", Last: "꾸밈:은", Pred: true, Modifiable: true}}}, true},
		{"before an opening sign", rule.Around{Self: apart, After: []rule.Near{{Stem: "“", Kind: "열기:1"}}}, true},
		{"onto an opening sign", rule.Around{Self: glued, Before: []rule.Near{{Stem: "“", Kind: "열기:1"}}, After: []rule.Near{noun}}, true},
		{"before a unit", rule.Around{Self: apart, After: []rule.Near{{Stem: "개", Kind: "열쇠:체언", Modifiable: true}}}, false},
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
