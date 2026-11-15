//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestCopulaEndRuleGraph — the graph of copulaEndRule evaluated by hand on rule.Around: the ending 야 stands at the sentence's end and right before a full stop or a question mark, and falls right before an exclamation mark, a comma, an ellipsis, a closing quotation mark or a word; 야 read as the vocative falls even at the end; the ending 라고 stands right before a full stop
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestCopulaEndRuleGraph(t *testing.T) {
	ya := rule.Near{Stem: "차", Kind: "열쇠:체언", Last: "맺음:야", Marks: "서술격:이(빠짐)+맺음:야"}
	sign := func(name string) rule.Near { return rule.Near{Last: name, Marks: name} }
	cases := []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"at the end", rule.Around{Self: ya}, true},
		{"before a full stop", rule.Around{Self: ya, After: []rule.Near{sign("마침표")}}, true},
		{"before a question mark", rule.Around{Self: ya, After: []rule.Near{sign("물음표")}}, true},
		{"before an exclamation mark", rule.Around{Self: ya, After: []rule.Near{sign("느낌표")}}, false},
		{"before a comma", rule.Around{Self: ya, After: []rule.Near{sign("쉼표")}}, false},
		{"before an ellipsis", rule.Around{Self: ya, After: []rule.Near{sign("줄임표")}}, false},
		{"before a closing quotation mark", rule.Around{Self: ya, After: []rule.Near{{Stem: "\"", Kind: "닫기:1"}}}, false},
		{"before a word", rule.Around{Self: ya, After: []rule.Near{{Stem: "가", Kind: "열쇠:동사", Last: "맺음:자", Pred: true}}}, false},
		{"before a bare noun", rule.Around{Self: ya, After: []rule.Near{{Stem: "책", Kind: "열쇠:체언"}}}, false},
		{"the vocative at the end", rule.Around{Self: rule.Near{Stem: "차", Last: "호격:야", Marks: "호격:야"}}, false},
		{"라고 before a full stop", rule.Around{Self: rule.Near{Stem: "책", Last: "맺음:라고", Marks: "서술격:이+맺음:라고"}, After: []rule.Near{sign("마침표")}}, true},
	}
	g := copulaEndRule().graph
	for _, c := range cases {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, c.around)
		if v, _ := evaluate(g, ctx, false); (v > 0) != c.want {
			t.Errorf("%s: %v, want %v", c.name, v, c.want)
		}
	}
}
