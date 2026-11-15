//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNumberQuoteRuleGraph — the graph of numberQuoteRule evaluated by hand on rule.Around: the particle 라고 or 이라고 on a number or an ordinal number stands; on a noun it falls; 라고 read as an ending on a number falls
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestNumberQuoteRuleGraph(t *testing.T) {
	cases := []struct {
		name string
		self rule.Near
		want bool
	}{
		{"on a number", rule.Near{Stem: "②", Kind: "수", Last: "부사격:라고", Marks: "부사격:라고"}, true},
		{"이라고 on a number", rule.Near{Stem: "③", Kind: "수", Last: "부사격:이라고", Marks: "부사격:이라고"}, true},
		{"on an ordinal number", rule.Near{Stem: "제3", Kind: "차례", Last: "부사격:라고", Marks: "단위:자+부사격:라고"}, true},
		{"on a noun", rule.Near{Stem: "책", Kind: "열쇠:체언", Last: "부사격:이라고", Marks: "부사격:이라고"}, false},
		{"an ending on a number", rule.Near{Stem: "②", Kind: "수", Last: "맺음:라고", Marks: "맺음:라고"}, false},
	}
	g := numberQuoteRule().graph
	for _, c := range cases {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, rule.Around{Self: c.self, After: []rule.Near{{Stem: "하", Kind: "열쇠:동사", Last: "맺음:ㄴ다", Pred: true}}})
		if v, _ := evaluate(g, ctx, false); (v > 0) != c.want {
			t.Errorf("%s: %v, want %v", c.name, v, c.want)
		}
	}
}
