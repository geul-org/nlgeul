//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestComplementRuleGraph — the graph of complementRule evaluated by hand on rule.Around: 보격 이 before 되/아니 stands; the same before another stem falls; before a question word with 이/가 and then 되 it stands, and the question word itself with a 이/가 word before it falls but with none stands; a question word with 를 or one followed by 먹 does not lift what is before it; the stem 누 of the contraction reading is no question word
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestComplementRuleGraph(t *testing.T) {
	comp := rule.Near{Stem: "반장", Last: "보격:이", Marks: "보격:이"}
	ask := func(last string) rule.Near { return rule.Near{Stem: "누구", Last: last, Marks: last} }
	become := rule.Near{Stem: "되", Pred: true}
	eat := rule.Near{Stem: "먹", Pred: true}
	cases := []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before 되", rule.Around{Self: comp, After: []rule.Near{become}}, true},
		{"before 아니", rule.Around{Self: comp, After: []rule.Near{{Stem: "아니", Pred: true}}}, true},
		{"before 먹", rule.Around{Self: comp, After: []rule.Near{eat}}, false},
		{"at the end", rule.Around{Self: comp}, false},
		{"before 누구가 되", rule.Around{Self: comp, After: []rule.Near{ask("주격:가"), become}}, true},
		{"before 누구가 먹", rule.Around{Self: comp, After: []rule.Near{ask("주격:가"), eat}}, false},
		{"before 누구를 되", rule.Around{Self: comp, After: []rule.Near{ask("목적격:를"), become}}, false},
		{"before 누구가 and the end", rule.Around{Self: comp, After: []rule.Near{ask("주격:가")}}, false},
		{"question word alone", rule.Around{Self: ask("보격:가"), After: []rule.Near{become}}, true},
		{"question word after a subject", rule.Around{Self: ask("보격:가"), Before: []rule.Near{{Stem: "반장", Last: "주격:이"}}, After: []rule.Near{become}}, false},
		{"question word after a complement", rule.Around{Self: ask("보격:가"), Before: []rule.Near{{Stem: "반장", Last: "보격:이"}}, After: []rule.Near{become}}, false},
		{"question word after an object", rule.Around{Self: ask("보격:가"), Before: []rule.Near{{Stem: "밥", Last: "목적격:을"}}, After: []rule.Near{become}}, true},
		{"stem 누 of the contraction", rule.Around{Self: rule.Near{Stem: "누", Last: "보격:가", Marks: "보격:가"}, Before: []rule.Near{{Stem: "반장", Last: "주격:이"}}, After: []rule.Near{become}}, true},
		{"a subject, not a complement", rule.Around{Self: rule.Near{Stem: "반장", Last: "주격:이", Marks: "주격:이"}, After: []rule.Near{become}}, false},
	}
	g := complementRule().graph
	for _, c := range cases {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, c.around)
		if v, _ := evaluate(g, ctx, false); (v > 0) != c.want {
			t.Errorf("%s: %v, want %v", c.name, v, c.want)
		}
	}
}
