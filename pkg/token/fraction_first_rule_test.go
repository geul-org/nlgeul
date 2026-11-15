//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestFractionFirstRule — the rule of a sentence is named as the trace prints it, adds bonusOrder (an order among names of the same letters) and waits for nothing, and speaks only of a candidate whose stem is named 나눔 (not a date of the same letters, not a number)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestFractionFirstRule(t *testing.T) {
	r := fractionFirstRule()
	if r.name != ruleFractionFirst || r.graph == nil || r.bonus != bonusOrder || r.waits != nil {
		t.Fatal(r.name, r.bonus, r.waits != nil)
	}
	cases := []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "1/12", Kind: "나눔"}, true},
		{rule.Near{Stem: "1/12", Kind: "나눔", Last: "부사격:에", Marks: "부사격:에"}, true},
		{rule.Near{Stem: "1/12", Kind: "날짜"}, false},
		{rule.Near{Stem: "12", Kind: "수"}, false},
		{rule.Near{}, false},
	}
	for _, c := range cases {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v; want %v", c.near, got, c.want)
		}
	}
}
