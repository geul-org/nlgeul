//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestNounUnmodifiedRuleGraph — the graph of nounUnmodifiedRule evaluated by hand on rule.Around: the noun 수 stands after a noun, at the head and after a particle, falls after a modifying closing; the noun 줄 falls right before 알다 or 모르다 even after a noun
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounUnmodifiedRuleGraph(t *testing.T) {
	su := rule.Near{Stem: "수", Kind: "열쇠:체언"}
	jul := rule.Near{Stem: "줄", Kind: "열쇠:체언"}
	noun := []rule.Near{{Stem: "인구", Kind: "열쇠:체언"}}
	r := nounUnmodifiedRule(adnBound)
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after a noun", rule.Around{Self: su, Before: noun}, true},
		{"at the head", rule.Around{Self: su}, true},
		{"after a particle", rule.Around{Self: su, Before: []rule.Near{{Stem: "사람", Kind: "열쇠:체언", Last: "관형격:의", Marks: "관형격:의"}}}, true},
		{"after a modifying closing", rule.Around{Self: su, Before: []rule.Near{{Stem: "알", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Pred: true, Marks: "꾸밈:ㄹ"}}}, false},
		{"줄 after a noun", rule.Around{Self: jul, Before: noun}, true},
		{"줄 before 알다", rule.Around{Self: jul, Before: noun, After: []rule.Near{{Stem: "알", Kind: "열쇠:동사", Last: "맺음:고", Pred: true}}}, false},
		{"줄 before 모르다", rule.Around{Self: jul, Before: noun, After: []rule.Near{{Stem: "모르", Kind: "열쇠:동사", Last: "맺음:다", Pred: true}}}, false},
		{"줄 before 걸다", rule.Around{Self: jul, Before: noun, After: []rule.Near{{Stem: "걸", Kind: "열쇠:동사", Last: "맺음:다", Pred: true}}}, true},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
