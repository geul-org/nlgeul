//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAdnHeadRuleGraph — the graph of adnHeadRule evaluated by hand on rule.Around: an adnominal is rebutted before a word no adnominal may stand before (an adverb, a predicate, a full stop); not before one an adnominal may (or one not marked so), nor at the end, nor before an auxiliary of a dependent noun and 하다, a word ending in -고, an adverb of the closed set 꾸밈 사이 부사, or a word followed by a modifying form; the case exception table's row of 각 lifts it for that word
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdnHeadRuleGraph(t *testing.T) {
	r := adnHeadRule()
	self := rule.Near{Stem: "이", Kind: "관형사:이", Text: "이"}
	unmod := func(stem, kind, last string) rule.Near {
		return rule.Near{Stem: stem, Kind: kind, Last: last, Marks: last, Unmodifiable: true}
	}
	adverb := unmod("역시", "열쇠:부사", "")
	for _, c := range []struct {
		name  string
		self  rule.Near
		after []rule.Near
		want  bool
	}{
		{"before an adverb", self, []rule.Near{adverb, unmod("어렵", "열쇠:형용사", "맺음:다")}, true},
		{"before a predicate", self, []rule.Near{unmod("가", "열쇠:동사", "맺음:다")}, true},
		{"before a full stop", rule.Near{Stem: "전", Kind: "열쇠:관형사", Text: "전"}, []rule.Near{{Last: "마침표", Unmodifiable: true}}, true},
		{"before a word an adnominal may modify", self, []rule.Near{{Stem: "책", Kind: "열쇠:체언"}}, false},
		{"before an adverb not marked unmodifiable", self, []rule.Near{{Stem: "잠시", Kind: "열쇠:부사"}}, false},
		{"at the end", self, nil, false},
		{"before 듯하다", rule.Near{Stem: "오", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Marks: "꾸밈:ㄹ"}, []rule.Near{unmod("듯하", "열쇠:형용사", "맺음:다")}, false},
		{"before -고", self, []rule.Near{unmod("맥없", "열쇠:형용사", "맺음:고")}, false},
		{"before 가장", self, []rule.Near{unmod("가장", "열쇠:부사", "")}, false},
		{"before a clause of adnominals", self, []rule.Near{adverb, {Stem: "크", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ"}}, false},
		{"각 of the case exception table", rule.Near{Stem: "각", Kind: "열쇠:관형사", Text: "각"}, []rule.Near{unmod("기각하", "열쇠:동사", "맺음:ㄴ다")}, false},
	} {
		if got := graphHolds(t, r, rule.Around{Self: c.self, After: c.after}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
