//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAdnHeadRule — the rule judged on the adnominal is named as the trace prints it, is a rebuttal of the force forceUngrammatical that adds nothing, speaks of and waits for an adnominal (a modifying form, a determiner standing alone) and of nothing else (a noun, a predicate with another closing, a determiner's letters with a particle)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdnHeadRule(t *testing.T) {
	r := adnHeadRule()
	if r.name != ruleAdnHead || r.force != forceUngrammatical || r.bonus != 0 || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "크", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ", Marks: "꾸밈:ㄴ", Pred: true}, true},
		{rule.Near{Stem: "이", Kind: "관형사:이"}, true},
		{rule.Near{Stem: "책", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "크", Kind: "열쇠:형용사", Last: "맺음:다", Marks: "맺음:다", Pred: true}, false},
		{rule.Near{Stem: "이", Kind: "관형사:이", Last: "보조사:는", Marks: "보조사:는"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
