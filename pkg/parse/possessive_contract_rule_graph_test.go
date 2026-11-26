//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPossessiveContractRuleGraph — the graph of possessiveContractRule evaluated by hand on rule.Around: 내 read as 나 with 의 stands, written apart, before a noun, a modifying form and an opening sign, after a word with a particle, and written onto an opening sign; it falls before an adverb, a predicate, a sign and at the end of the sentence, written onto a word, and right after a nominal with no particle
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPossessiveContractRuleGraph(t *testing.T) {
	apart := rule.Near{Stem: "나", Kind: "열쇠:체언", Last: "관형격:의", Marks: "관형격:의", Text: "내"}
	glued := apart
	glued.Glued = true
	noun := rule.Near{Stem: "말", Kind: "열쇠:체언", Last: "목적격:을", Marks: "목적격:을", Modifiable: true}
	r := possessiveContractRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before a noun", rule.Around{Self: apart, After: []rule.Near{noun}}, true},
		{"before a modifying form", rule.Around{Self: apart, After: []rule.Near{{Stem: "작", Kind: "열쇠:형용사", Last: "꾸밈:은", Pred: true, Modifiable: true}}}, true},
		{"before an opening sign", rule.Around{Self: apart, After: []rule.Near{{Stem: "“", Kind: "열기:1"}}}, true},
		{"after a word with a particle", rule.Around{Self: apart, Before: []rule.Near{{Stem: "이것", Kind: "열쇠:체언", Last: "보조사:은", Marks: "보조사:은"}}, After: []rule.Near{noun}}, true},
		{"onto an opening sign", rule.Around{Self: glued, Before: []rule.Near{{Stem: "“", Kind: "열기:1"}}, After: []rule.Near{noun}}, true},
		{"before an adverb", rule.Around{Self: apart, After: []rule.Near{{Stem: "다시", Kind: "열쇠:부사"}, noun}}, false},
		{"before a predicate", rule.Around{Self: apart, After: []rule.Near{{Stem: "버리", Kind: "열쇠:동사", Last: "맺음:다", Pred: true}}}, false},
		{"before a sign", rule.Around{Self: apart, After: []rule.Near{{Stem: ",", Kind: "쉼표", Glued: true}}}, false},
		{"at the end", rule.Around{Self: apart}, false},
		{"onto a word", rule.Around{Self: glued, Before: []rule.Near{{Stem: "학교", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에"}}, After: []rule.Near{noun}}, false},
		{"after a nominal with no particle", rule.Around{Self: apart, Before: []rule.Near{{Stem: "기간", Kind: "열쇠:체언", Nominal: true}}, After: []rule.Near{noun}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
