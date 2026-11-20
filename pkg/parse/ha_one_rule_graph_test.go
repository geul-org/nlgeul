//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestHaOneRuleGraph — the graph of haOneRule evaluated by hand on rule.Around: 하다 with ㄴ written apart stands after the object particle, after 로, after the dative 에게 and after the closing -게, and falls after the nominative, after an adverb and at the head of a sentence, and before a unit a numeral determiner counts (일을 한 번) but not before 자, which a modifying form stands before as well (일을 한 자); written onto a noun it stands, and written onto an opening sign or onto the modifying 는 it falls
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestHaOneRuleGraph(t *testing.T) {
	apart := rule.Near{Text: "한", Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Pred: true}
	glued := apart
	glued.Glued = true
	noun := rule.Near{Stem: "사람", Kind: "열쇠:체언", Modifiable: true}
	with := func(last string) rule.Near {
		return rule.Near{Stem: "일", Kind: "열쇠:체언", Last: last, Marks: last, Nominal: true}
	}
	r := haOneRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after 을", rule.Around{Self: apart, Before: []rule.Near{with("목적격:을")}, After: []rule.Near{noun}}, true},
		{"after 로", rule.Around{Self: apart, Before: []rule.Near{with("부사격:로")}, After: []rule.Near{noun}}, true},
		{"after 에게", rule.Around{Self: apart, Before: []rule.Near{with("부사격:에게")}, After: []rule.Near{noun}}, true},
		{"after -게", rule.Around{Self: apart, Before: []rule.Near{{Stem: "오", Kind: "열쇠:동사", Last: "맺음:게", Marks: "맺음:게", Pred: true}}, After: []rule.Near{noun}}, true},
		{"after 을, before 번", rule.Around{Self: apart, Before: []rule.Near{with("목적격:을")}, After: []rule.Near{{Stem: "번", Kind: "열쇠:체언", Modifiable: true}}}, false},
		{"after 을, before 자", rule.Around{Self: apart, Before: []rule.Near{with("목적격:을")}, After: []rule.Near{{Stem: "자", Kind: "열쇠:체언", Modifiable: true}}}, true},
		{"after 이", rule.Around{Self: apart, Before: []rule.Near{with("주격:이")}, After: []rule.Near{noun}}, false},
		{"after an adverb", rule.Around{Self: apart, Before: []rule.Near{{Stem: "또", Kind: "열쇠:부사"}}, After: []rule.Near{noun}}, false},
		{"at the head", rule.Around{Self: apart, After: []rule.Near{noun}}, false},
		{"onto a noun", rule.Around{Self: glued, Before: []rule.Near{{Stem: "창백", Kind: "열쇠:체언", Nominal: true}}, After: []rule.Near{noun}}, true},
		{"onto an opening sign", rule.Around{Self: glued, Before: []rule.Near{{Stem: "“", Kind: "열기:1"}}, After: []rule.Near{noun}}, false},
		{"onto the modifying 는", rule.Around{Self: glued, Before: []rule.Near{{Stem: "있", Kind: "보조 용언:있", Last: "꾸밈:는", Marks: "꾸밈:는", Pred: true}}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
