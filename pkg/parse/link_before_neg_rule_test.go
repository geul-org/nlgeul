//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestLinkBeforeNegRule — the rule is named as the trace prints it, adds bonusShortTaken, and speaks of and waits for a candidate of a predicate's stem (a verb, an adjective, an auxiliary verb) with the ending 지 or 진 among its markers (not a noun, not a new word, not a noun with the copula dropped and 지, not a verb with another ending)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLinkBeforeNegRule(t *testing.T) {
	r := linkBeforeNegRule()
	if r.name != ruleLinkBeforeNeg || r.bonus != bonusShortTaken || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:지", Marks: "맺음:지", Pred: true}, true},
		{rule.Near{Stem: "적", Kind: "열쇠:형용사", Last: "보조사:도", Marks: "맺음:지+보조사:도", Pred: true}, true},
		{rule.Near{Stem: "타", Kind: "열쇠:동사", Last: "맺음:진(지는)", Marks: "맺음:진(지는)", Pred: true}, true},
		{rule.Near{Stem: "주", Kind: "보조 용언:주", Last: "맺음:지", Marks: "맺음:지", Pred: true}, true},
		{rule.Near{Stem: "가지", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "아버", Kind: "", Last: "맺음:지", Marks: "맺음:지", Pred: true}, false},
		{rule.Near{Stem: "가", Kind: "열쇠:체언", Last: "맺음:지", Marks: "서술격:이(빠짐)+맺음:지", Pred: true}, false},
		{rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:고", Marks: "맺음:고", Pred: true}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
