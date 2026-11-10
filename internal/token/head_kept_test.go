//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestHeadKept — a path's head, an option's dashes and a sign before Latin letters are kept; a dash before Hangul or a digit, a sign before Hangul, an unlisted run of signs, and a core with nothing before it are not
package token

import "testing"

func TestHeadKept(t *testing.T) {
	cases := map[string]int{"~/.a/b": 0, "(./a": 1, "/a": 0, "--w": 0, "-v": 0, "-가": 1, "-3": 1, "@a": 0, "#가": 1, "$A": 0, "._a": 1, "~~a": 2, "a": 0, "(a": 1}
	for in, want := range cases {
		all := []rune(in)
		from := 0
		for from < len(all) && all[from] != 'a' && all[from] != 'w' && all[from] != 'v' && all[from] != 'A' && all[from] != '가' && all[from] != '3' {
			from++
		}
		if got := headKept(all, from); got != want {
			t.Errorf("%s: %d, want %d", in, got, want)
		}
	}
}
