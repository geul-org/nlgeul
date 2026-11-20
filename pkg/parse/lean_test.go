//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestLean — a quotation mark written the same way at both ends leans to opening with a space or the sentence's start on its left and a letter or digit on its right, to closing with something other than a space on its left and a space or the sentence's end on its right, and to neither between two letters, between two spaces or with a sign on its right after a space; a curly quotation mark and a bracket lean to nothing
package parse

import "testing"

func TestLean(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text string
		at   int
		want string
	}{
		{"'봄", 0, openName},
		{"가 '봄", 2, openName},
		{"가 \"3", 2, openName},
		{"간다'", 2, shutName},
		{"간다' 가", 2, shutName},
		{"다.\" 가", 2, shutName},
		{"가'나", 1, ""},
		{"가 ' 나", 2, ""},
		{"', 가", 0, ""},
		{"'", 0, ""},
		{"‘봄", 0, ""},
		{"가 (봄", 2, ""},
		{"봄) 가", 1, ""},
	} {
		cut := newCutter(ix, c.text, false)
		if got := cut.lean(c.at, pairRows[cut.r[c.at]]); got != c.want {
			t.Errorf("%q at %d: %q, want %q", c.text, c.at, got, c.want)
		}
	}
}
