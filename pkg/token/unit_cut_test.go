//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestUnitCut — a run of digits and a Latin unit of the table is cut where the unit begins; no cut for a unit alone, digits alone, a word that is no unit of the table, or digits and a Hangul unit
package token

import "testing"

func TestUnitCut(t *testing.T) {
	for _, c := range []struct {
		text string
		cut  int
		ok   bool
	}{
		{"12ms", 2, true},
		{"007kg", 3, true},
		{"ms", 0, false},
		{"12", 2, false},
		{"12xq", 2, false},
		{"12개", 2, false},
	} {
		s := []rune(c.text)
		cut, ok := unitCut(s, [2]int{0, len(s)})
		if ok != c.ok || (ok && cut != c.cut) {
			t.Errorf("%q: %d %v, want %d %v", c.text, cut, ok, c.cut, c.ok)
		}
	}
	s := []rune("가 5ms")
	if cut, ok := unitCut(s, [2]int{2, 5}); !ok || cut != 3 {
		t.Fatal("a run that begins inside the text", cut, ok)
	}
}
