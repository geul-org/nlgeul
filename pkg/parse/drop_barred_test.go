//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestDropBarred — right after a sign that closes a pair (a closing bracket, a quotation mark) a dropped copula stem is not barred; right after a sign that only opens one it is barred before 다 and not before 야; after a Latin letter, a digit, a unit's sign or a Hangul syllable nothing is barred
package parse

import "testing"

func TestDropBarred(t *testing.T) {
	for _, c := range []struct {
		prev  rune
		after string
		want  bool
	}{
		{')', "맺음:다", false}, {'\'', "맺음:냐고", false}, {'”', "맺음:다", false},
		{'(', "맺음:다", true}, {'“', "맺음:다", true}, {'c', "맺음:다", false}, {'3', "맺음:다", false}, {'%', "맺음:면", false},
		{'(', "맺음:야", false}, {'c', "맺음:야", false}, {'교', "맺음:다", false},
	} {
		if got := dropBarred(c.prev, c.after); got != c.want {
			t.Errorf("dropBarred(%q, %s) = %v", c.prev, c.after, got)
		}
	}
}
