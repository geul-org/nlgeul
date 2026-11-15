//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestDateFirstRule — the rule of a sentence is named as the trace prints it, adds bonusOrder (an order among names of the same letters) and waits for nothing, and speaks only of a candidate whose stem is named 날짜 and written with full stops (not a date written with a hyphen or a slash, not a version or a decimal point of the same letters)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDateFirstRule(t *testing.T) {
	r := dateFirstRule()
	if r.name != ruleDateFirst || r.graph == nil || r.bonus != bonusOrder || r.waits != nil {
		t.Fatal(r.name, r.bonus, r.waits != nil)
	}
	cases := []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "1989.8.25", Kind: "날짜"}, true},
		{rule.Near{Stem: "1985.10.16", Kind: "날짜", Last: "부사격:에", Marks: "마침표+부사격:에"}, true},
		{rule.Near{Stem: "1989.8.25", Kind: "버전"}, false},
		{rule.Near{Stem: "2021-5-27", Kind: "날짜"}, false},
		{rule.Near{Stem: "1/12", Kind: "날짜"}, false},
		{rule.Near{Stem: "4.7", Kind: "소수점"}, false},
		{rule.Near{}, false},
	}
	for _, c := range cases {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v; want %v", c.near, got, c.want)
		}
	}
}
