//ff:func feature=token type=helper control=sequence
//ff:what TestTwos — the places of a sign pair two by two, the odd last one left out; places to skip and an apostrophe inside a word do not count
package token

import "testing"

func TestTwos(t *testing.T) {
	got := twos([]rune("'가' '나"), '\'', nil)
	if len(got) != 2 || got[0] != 2 || got[2] != 0 {
		t.Fatal(got)
	}
	if got = twos([]rune("it's '가'"), '\'', nil); len(got) != 2 || got[5] != 7 {
		t.Fatal("an apostrophe inside a word", got)
	}
	if got = twos([]rune("\"가\""), '"', []bool{true, false, false}); len(got) != 0 {
		t.Fatal("a place to skip", got)
	}
}
