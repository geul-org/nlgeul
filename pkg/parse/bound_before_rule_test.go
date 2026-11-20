//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestBoundBeforeRule — the rule is named as the trace prints it, adds bonusShortTaken, and speaks of and waits for a candidate whose last piece is a modifying closing alone (not of a stem alone, a linking closing, a particle or a naming closing)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundBeforeRule(t *testing.T) {
	r := boundBeforeRule(boundNext, boundNominal)
	if r.name != ruleBoundBefore || r.bonus != bonusShortTaken || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "알", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Pred: true, Marks: "꾸밈:ㄹ"}, true},
		{rule.Near{Stem: "안", Kind: "열쇠:동사", Last: "꾸밈:은", Pred: true, Marks: "꾸밈:은"}, true},
		{rule.Near{Stem: "알", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "알", Kind: "열쇠:동사", Last: "맺음:고", Pred: true, Marks: "맺음:고"}, false},
		{rule.Near{Stem: "안", Kind: "열쇠:체언", Last: "보조사:은", Marks: "보조사:은"}, false},
		{rule.Near{Stem: "알", Kind: "열쇠:동사", Last: "이름:ㅁ", Pred: true, Marks: "이름:ㅁ"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
