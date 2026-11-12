//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNameOf — a name with a colon is a kind and a form, one with none a kind alone, the form keeps a colon of its own, and no name gives the empty reading
package token

import "testing"

func TestNameOf(t *testing.T) {
	cases := map[string]Name{
		"주격:가":    {"주격", "가"},
		"열쇠:체언":   {"열쇠", "체언"},
		"수":       {"수", ""},
		"시각::":    {"시각", ":"},
		"보조 용언:두": {"보조 용언", "두"},
		"":        {},
	}
	for in, want := range cases {
		if got := nameOf(in); got != want {
			t.Errorf("nameOf(%q) = %v, want %v", in, got, want)
		}
	}
}
