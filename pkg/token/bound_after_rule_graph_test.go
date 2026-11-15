//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestBoundAfterRuleGraph — the graph of boundAfterRule evaluated by hand on rule.Around: a dependent noun stands after a modifying closing (수 after 갈, 거 after 하는), falls after a noun, at the head and after a linking closing, and falls when it stands after a nominal as well (시, 때) even after a modifying closing
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundAfterRuleGraph(t *testing.T) {
	adn := []rule.Near{{Stem: "가", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Pred: true, Marks: "꾸밈:ㄹ"}}
	bound := func(stem string) rule.Near { return rule.Near{Stem: stem, Kind: "의존명사:" + stem} }
	r := boundAfterRule(boundNominal)
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"수 after 갈", rule.Around{Self: bound("수"), Before: adn}, true},
		{"거 after 하는", rule.Around{Self: bound("거"), Before: []rule.Near{{Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:는", Pred: true, Marks: "꾸밈:는"}}}, true},
		{"after a noun", rule.Around{Self: bound("수"), Before: []rule.Near{{Stem: "인구", Kind: "열쇠:체언"}}}, false},
		{"at the head", rule.Around{Self: bound("수")}, false},
		{"after a linking closing", rule.Around{Self: bound("수"), Before: []rule.Near{{Stem: "가", Kind: "열쇠:동사", Last: "맺음:고", Pred: true, Marks: "맺음:고"}}}, false},
		{"시 after a modifying closing", rule.Around{Self: bound("시"), Before: adn}, false},
		{"때 after a modifying closing", rule.Around{Self: bound("때"), Before: adn}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
