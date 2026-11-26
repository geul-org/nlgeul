//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCutterQuoted — a marker that takes a quotation is written right after the sign when the Hangul syllables after it are one of quoteMarkers and no letter or digit follows them ("라고 and a space, 다'는 at the end, ”고, before a comma); not when the syllables are another word (는데, 이), when a letter or digit follows them, or when nothing or no Hangul stands after the sign
package parse

import "testing"

func TestCutterQuoted(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text string
		at   int
		want bool
	}{
		{"\"라고 했다", 0, true},
		{"좋다'는", 2, true},
		{"좋다”고, 또", 2, true},
		{"곰‘이라고 한다", 1, true},
		{"곰\"란", 1, true},
		{"좋다\"는데", 2, false},
		{"곰'이 왔다", 1, false},
		{"곰\"고A", 1, false},
		{"곰\"고3", 1, false},
		{"곰\"", 1, false},
		{"곰\" 고", 1, false},
		{"\"A라고", 0, false},
	} {
		if got := newCutter(ix, c.text, false).quoted(c.at); got != c.want {
			t.Errorf("%q at %d: %v", c.text, c.at, got)
		}
	}
}
