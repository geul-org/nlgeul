//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNegKindRuleGraph — the graphs of negKindRule evaluated by hand on rule.Around, both ways as the index puts them: the adjective 못하 is rebutted after the -지 form of a verb (also with a particle after 지, 진, and after an auxiliary verb of the closed set), not after that of an adjective, not after another ending, not with nothing before it; the verb 못하 is rebutted after the -지 form of an adjective and not after that of a verb
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNegKindRuleGraph(t *testing.T) {
	adj := negKindRule("열쇠:형용사", append([]string{"열쇠:동사"}, auxNamesOf([]string{"주", "있"})...), whyNotVerbBefore)
	verb := negKindRule("열쇠:동사", []string{"열쇠:형용사"}, whyNotAdjBefore)
	self := rule.Near{Stem: "못하", Kind: "열쇠:형용사", Last: "맺음:다", Marks: "맺음:다", Pred: true}
	before := func(stem, kind, marks string) []rule.Near {
		return []rule.Near{{Stem: stem, Kind: kind, Marks: marks, Pred: true}}
	}
	for _, c := range []struct {
		name   string
		r      sentenceRule
		before []rule.Near
		want   bool
	}{
		{"adjective after a verb's -지", adj, before("가", "열쇠:동사", "맺음:지"), true},
		{"adjective after -지도", adj, before("가", "열쇠:동사", "맺음:지+보조사:도"), true},
		{"adjective after 진", adj, before("가", "열쇠:동사", "맺음:진(지는)"), true},
		{"adjective after an auxiliary verb's -지", adj, before("주", "보조 용언:주", "맺음:지"), true},
		{"adjective after an adjective's -지", adj, before("편하", "열쇠:형용사", "맺음:지"), false},
		{"adjective after a verb's -고", adj, before("가", "열쇠:동사", "맺음:고"), false},
		{"adjective after a new word with 지", adj, before("쪼간", "", "맺음:지"), false},
		{"adjective with nothing before", adj, nil, false},
		{"verb after an adjective's -지", verb, before("편하", "열쇠:형용사", "맺음:지"), true},
		{"verb after a verb's -지", verb, before("가", "열쇠:동사", "맺음:지"), false},
		{"verb after an auxiliary verb's -지", verb, before("주", "보조 용언:주", "맺음:지"), false},
	} {
		if got := graphHolds(t, c.r, rule.Around{Self: self, Before: c.before}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
