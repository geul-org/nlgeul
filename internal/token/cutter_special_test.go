//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCutterSpecial — a place stands inside a special stretch (a path, an address) from where the unit's letters begin up to the next space, and not past its last letter that is no Hangul syllable
package token

import "testing"

func TestCutterSpecial(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text string
		at   int
		want bool
	}{
		{"/a/b+c", 4, true},
		{"/a/b+c 가", 4, true},
		{"곰+돌", 1, false},
		{"가 /a/b+c", 6, false},
	} {
		if got := newCutter(ix, c.text, false).special(c.at); got != c.want {
			t.Errorf("%q at %d: %v", c.text, c.at, got)
		}
	}
	c := newCutter(ix, "가 /a/b+c", false)
	c.at = 2
	if !c.special(6) {
		t.Fatal("from where the unit begins")
	}
}
