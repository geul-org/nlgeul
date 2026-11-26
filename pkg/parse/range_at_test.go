//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestRangeAt — the tilde written right after a number and its unit and right before a digit is the range sign of the number's word (3월~5월, 3개∼5개); not a tilde after a unit with no number (원~5), before a letter that is no digit (3층~옥상), nor a sign that is no tilde (3월/5월)
package parse

import "testing"

func TestRangeAt(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		text string
		q    int
		want bool
	}{
		{"3월~5월", 2, true},
		{"3개∼5개", 2, true},
		{"원~5", 1, false},
		{"3층~옥상", 2, false},
		{"3월/5월", 2, false},
	} {
		sc := &scan{ix: r.ix, s: []rune(c.text)}
		sc.prepare()
		if got := sc.rangeAt(c.q); got != c.want {
			t.Errorf("%q at %d: %v, want %v", c.text, c.q, got, c.want)
		}
	}
}
