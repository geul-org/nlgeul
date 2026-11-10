//ff:func feature=token type=reader control=sequence
//ff:what TestEndsOf — the word ends of each reading joined with ",", the readings with ";"; no readings give an empty text
package token

import "testing"

func TestEndsOf(t *testing.T) {
	first := &path{prev: &path{}, last: step{end: 2}}
	two := &path{prev: first, last: step{end: 5}}
	if got := endsOf([]*path{two, first}); got != "2,5;2" {
		t.Fatal(got)
	}
	if endsOf(nil) != "" {
		t.Fatal("empty")
	}
}
