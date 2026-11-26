//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPossessiveContractRule — the rule is named as the trace prints it, adds bonusOverCount, and speaks of and waits for the candidate that reads the word 내 or 네 as 나 or 너 with the particle 의 fused; not 나의 written out, not 내 read as the nominal with no particle, not 내 with another particle, not the verb 내
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPossessiveContractRule(t *testing.T) {
	r := possessiveContractRule()
	if r.name != rulePossessiveContract || r.bonus != bonusOverCount || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "나", Kind: "열쇠:체언", Last: "관형격:의", Marks: "관형격:의", Text: "내"}, true},
		{rule.Near{Stem: "나", Kind: "열쇠:체언", Last: "관형격:의", Marks: "관형격:의", Text: "내", Glued: true}, true},
		{rule.Near{Stem: "너", Kind: "열쇠:체언", Last: "관형격:의", Marks: "관형격:의", Text: "네"}, true},
		{rule.Near{Stem: "나", Kind: "열쇠:체언", Last: "관형격:의", Marks: "관형격:의", Text: "나의"}, false},
		{rule.Near{Stem: "네", Kind: "열쇠:체언", Text: "네"}, false},
		{rule.Near{Stem: "네", Kind: "열쇠:홀로", Text: "네"}, false},
		{rule.Near{Stem: "네", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가", Text: "네가"}, false},
		{rule.Near{Stem: "내", Kind: "열쇠:체언", Text: "내"}, false},
		{rule.Near{Stem: "내", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가", Text: "내가"}, false},
		{rule.Near{Stem: "내", Kind: "열쇠:동사", Last: "맺음:어(합침)", Marks: "맺음:어(합침)", Pred: true, Text: "내"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
