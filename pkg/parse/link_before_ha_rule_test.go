//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestLinkBeforeHaRule — the rule is named as the trace prints it, adds bonusOverCount, waits for the word after it, and speaks of a predicate's candidate with the linking ending -어야 (written or merged, as an auxiliary verb's too), not of a noun 내야, a noun with the particle 야, a predicate with another closing, nor a new word with 어야
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLinkBeforeHaRule(t *testing.T) {
	r := linkBeforeHaRule()
	if r.name != ruleLinkBeforeHa || r.bonus != bonusOverCount || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "내", Kind: "열쇠:동사", Last: "맺음:어야(합침)", Marks: "맺음:어야(합침)", Pred: true}, true},
		{rule.Near{Stem: "먹", Kind: "열쇠:동사", Last: "맺음:어야", Marks: "맺음:어야", Pred: true}, true},
		{rule.Near{Stem: "좋", Kind: "열쇠:형용사", Last: "맺음:아야", Marks: "맺음:아야", Pred: true}, true},
		{rule.Near{Stem: "주", Kind: "보조 용언:주", Last: "맺음:어야(합침)", Marks: "맺음:어야(합침)", Pred: true}, true},
		{rule.Near{Stem: "내야", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "자", Kind: "열쇠:체언", Last: "보조사:야", Marks: "보조사:야"}, false},
		{rule.Near{Stem: "내", Kind: "열쇠:동사", Last: "맺음:어도(합침)", Marks: "맺음:어도(합침)", Pred: true}, false},
		{rule.Near{Stem: "드러", Kind: "", Last: "맺음:어야", Marks: "맺음:어야", Pred: true}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
