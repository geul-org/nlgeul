//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestPriorNounRuleGraph — the graph of priorNounRule evaluated by hand on rule.Around: 전 the noun stands right after a -기 or -음 form and right after a number with its unit or an approximation, and falls after a noun, after a modifying form, after a bare number and at the head of the sentence
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPriorNounRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "전", Kind: "열쇠:체언"}
	before := func(stem, kind, last string) []rule.Near {
		return []rule.Near{{Stem: stem, Kind: kind, Last: last, Marks: last}}
	}
	r := priorNounRule()
	for _, c := range []struct {
		name   string
		before []rule.Near
		want   bool
	}{
		{"after -기", before("먹", "열쇠:동사", "이름:기"), true},
		{"after -음", before("죽", "열쇠:동사", "이름:음"), true},
		{"after a unit", before("3", "수", "단위:년"), true},
		{"after an approximation", before("10", "수", "어림:여"), true},
		{"after a noun", before("사람", "열쇠:체언", ""), false},
		{"after a modifying form", before("먹", "열쇠:동사", "꾸밈:은"), false},
		{"after a bare number", before("3", "수", ""), false},
		{"at the head", nil, false},
	} {
		if got := graphHolds(t, r, rule.Around{Self: self, Before: c.before}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
