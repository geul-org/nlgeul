//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestBoundAfterRuleAdverbial — after a modifying closing the graph of boundAfterRule falls for an adverbial dependent noun with the subject's or the complement's 이/가 (듯+이, 채+가) and stands for one standing bare or with another particle (듯, 듯+도, 채+로) and for a dependent noun that is not adverbial with 이/가 (수+가, 것+이)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundAfterRuleAdverbial(t *testing.T) {
	adn := []rule.Near{{Stem: "미안하", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ", Pred: true, Marks: "꾸밈:ㄴ"}}
	bound := func(stem, marks string) rule.Near {
		n := rule.Near{Stem: stem, Kind: "의존명사:" + stem, Marks: marks}
		if marks != "" {
			n.Last = marks
		}
		return n
	}
	r := boundAfterRule(boundNominal)
	for _, c := range []struct {
		name string
		self rule.Near
		want bool
	}{
		{"듯+이 (subject)", bound("듯", "주격:이"), false},
		{"듯+이 (complement)", bound("듯", "보격:이"), false},
		{"채+가", bound("채", "주격:가"), false},
		{"듯", bound("듯", ""), true},
		{"듯+도", bound("듯", "보조사:도"), true},
		{"채+로", bound("채", "부사격:로"), true},
		{"수+가", bound("수", "주격:가"), true},
		{"것+이", bound("것", "주격:이"), true},
	} {
		if got := graphHolds(t, r, rule.Around{Self: c.self, Before: adn}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
