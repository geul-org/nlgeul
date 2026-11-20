//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAdnHeadAfterRuleGraph — the graph of adnHeadAfterRule evaluated by hand on rule.Around: a candidate no adnominal modifies is rebutted right after a word each of whose candidates is an adnominal; not after one that could be read otherwise, nor with nothing before it; not when its stem is an auxiliary of a dependent noun and 하다, an adverb of the closed set 꾸밈 사이 부사, or a negative adverb
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdnHeadAfterRuleGraph(t *testing.T) {
	r := adnHeadAfterRule()
	firm := []rule.Near{{Stem: "스무", Kind: "열쇠:관형사", AdnominalAll: true}}
	loose := []rule.Near{{Stem: "이", Kind: "관형사:이"}}
	for _, c := range []struct {
		name   string
		self   rule.Near
		before []rule.Near
		want   bool
	}{
		{"an adverb after a determiner alone", rule.Near{Stem: "채", Kind: "열쇠:부사"}, firm, true},
		{"a predicate after a determiner alone", rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:다", Marks: "맺음:다"}, firm, true},
		{"after a word read otherwise as well", rule.Near{Stem: "채", Kind: "열쇠:부사"}, loose, false},
		{"with nothing before", rule.Near{Stem: "채", Kind: "열쇠:부사"}, nil, false},
		{"듯하다", rule.Near{Stem: "듯하", Kind: "열쇠:형용사", Last: "맺음:다", Marks: "맺음:다"}, firm, false},
		{"가장", rule.Near{Stem: "가장", Kind: "열쇠:부사"}, firm, false},
		{"안", rule.Near{Stem: "안", Kind: "열쇠:부사"}, firm, false},
	} {
		if got := graphHolds(t, r, rule.Around{Self: c.self, Before: c.before}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
