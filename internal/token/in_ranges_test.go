//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestInRanges — a letter at either end of a range or between is in it, one just outside is not, and nothing is in no range
package token

import "testing"

func TestInRanges(t *testing.T) {
	rs := []runeRange{{'a', 'c'}, {'x', 'z'}}
	for r, want := range map[rune]bool{'a': true, 'b': true, 'c': true, 'd': false, 'w': false, 'x': true, 'z': true, '{': false} {
		if got := inRanges(r, rs); got != want {
			t.Fatalf("%q: %v", r, got)
		}
	}
	if inRanges('a', nil) {
		t.Fatal("no range")
	}
}
