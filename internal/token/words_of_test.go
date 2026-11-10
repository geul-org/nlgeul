//ff:func feature=token type=reader control=sequence
//ff:what TestWordsOf — the words of a reading in order; none for the empty reading
package token

import "testing"

func TestWordsOf(t *testing.T) {
	p := grow(grow(&path{verdict: 1}, eojeol{end: 2, marked: "a"}), eojeol{end: 5, marked: "b"})
	got := wordsOf(p)
	if len(got) != 2 || got[0].end != 2 || got[1].end != 5 || wordsOf(&path{}) != nil {
		t.Fatal(got)
	}
}
