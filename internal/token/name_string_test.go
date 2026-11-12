//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNameString — a candidate prints as kind:form, a kind alone without a colon, a form that is a colon after the kind's colon, and nothing for the empty candidate
package token

import "testing"

func TestNameString(t *testing.T) {
	cases := map[Name]string{
		{"주격", "가"}:  "주격:가",
		{"열쇠", "체언"}: "열쇠:체언",
		{"소수점", ""}:  "소수점",
		{"시각", ":"}:  "시각::",
		{}:           "",
	}
	for in, want := range cases {
		if got := in.String(); got != want {
			t.Errorf("%#v prints %q, want %q", in, got, want)
		}
	}
}
