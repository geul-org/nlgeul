//ff:func feature=token type=reader control=sequence
//ff:what TestWordsOf — the words of a candidate in order; none for the empty candidate
package token

import "testing"

func TestWordsOf(t *testing.T) {
	p := grow(grow(&path{verdict: 1}, step{end: 2, sig: "a"}), step{end: 5, sig: "b"})
	got := wordsOf(p)
	if len(got) != 2 || got[0].end != 2 || got[1].end != 5 || wordsOf(&path{}) != nil {
		t.Fatal(got)
	}
}
