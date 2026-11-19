//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAdnHeadAfterRule — the rule judged on the word after the adnominal is named as adnHeadRule is, is a rebuttal of the force forceUngrammatical that adds nothing and reads nothing after the candidate, and speaks of a candidate no adnominal modifies that does not end in a linking closing (an adverb, a predicate with a final closing, an interjection), not of one that does (-어, -고), a nominal or a modifying form
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdnHeadAfterRule(t *testing.T) {
	r := adnHeadAfterRule()
	if r.name != ruleAdnHead || r.force != forceUngrammatical || r.bonus != 0 || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "채", Kind: "열쇠:부사"}, true},
		{rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:ㄴ다", Marks: "맺음:ㄴ다", Pred: true}, true},
		{rule.Near{Stem: "아", Kind: "열쇠:홀로"}, true},
		{rule.Near{Stem: "남", Kind: "열쇠:동사", Last: "맺음:아", Marks: "맺음:아", Pred: true}, false},
		{rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:고", Marks: "맺음:고", Pred: true}, false},
		{rule.Near{Stem: "채", Kind: "의존명사:채"}, false},
		{rule.Near{Stem: "크", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ", Marks: "꾸밈:ㄴ", Pred: true}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v, want %v", c.near, got, c.want)
		}
	}
}
