//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSignsAfter — the letters of a unit after the letters that are read: the signs trimmed from its end, none when nothing was trimmed there, those after the core of a unit with a head, and nothing for a unit with nothing to read
package token

import "testing"

func TestSignsAfter(t *testing.T) {
	cases := []struct {
		u    unit
		want string
	}{
		{unit{text: "문제야."}, "."},
		{unit{text: "문제야?!"}, "?!"},
		{unit{text: "문제야"}, ""},
		{unit{text: "(문제야..."}, "..."},
		{unit{text: ")를,", head: 1}, ","},
		{unit{text: "..."}, ""},
	}
	for _, c := range cases {
		sc := &scan{s: coreOf(c.u), head: c.u.head}
		if got := sc.signsAfter(c.u.text); got != c.want {
			t.Errorf("signsAfter(%q) = %q, want %q", c.u.text, got, c.want)
		}
	}
}
