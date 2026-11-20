//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestCopulaCommaRuleGraph — the graph of copulaCommaRule evaluated by hand on rule.Around: the copula 이 with the ending 나 stands right before a comma, and falls right before a full stop, before a word and at the end; the auxiliary particle 이나 falls before a comma (that the last piece is the ending 나 is asked by the rule's about, not by the graph)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestCopulaCommaRuleGraph(t *testing.T) {
	na := rule.Near{Stem: "것", Kind: "의존명사:것", Last: "맺음:나", Marks: "서술격:이+맺음:나"}
	comma := rule.Near{Last: "쉼표", Marks: "쉼표"}
	cases := []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before a comma", rule.Around{Self: na, After: []rule.Near{comma}}, true},
		{"before a full stop", rule.Around{Self: na, After: []rule.Near{{Last: "마침표", Marks: "마침표"}}}, false},
		{"before a word", rule.Around{Self: na, After: []rule.Near{{Stem: "책", Kind: "열쇠:체언"}}}, false},
		{"at the end", rule.Around{Self: na}, false},
		{"the particle 이나", rule.Around{Self: rule.Near{Stem: "것", Last: "보조사:이나", Marks: "보조사:이나"}, After: []rule.Near{comma}}, false},
	}
	g := copulaCommaRule().graph
	for _, c := range cases {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, c.around)
		if v, _ := evaluate(g, ctx, false); (v > 0) != c.want {
			t.Errorf("%s: %v, want %v", c.name, v, c.want)
		}
	}
}
