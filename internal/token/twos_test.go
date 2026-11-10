//ff:func feature=token type=helper control=sequence
//ff:what TestTwos — the places of a sign pair two by two, the odd last one left out; places to skip and an apostrophe between Latin letters or digits do not count (it's, and '90s alone has no partner); a closing quote with a particle written onto it pairs with the one opened a word before
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
	if got = twos([]rune("곰이 '사자 호랑이'는 싫다"), '\'', nil); len(got) != 2 || got[3] != 10 || got[10] != 3 {
		t.Fatal("a closing quote with a particle after it", got)
	}
	if got = twos([]rune("it's"), '\'', nil); len(got) != 0 {
		t.Fatal("it's", got)
	}
	if got = twos([]rune("'90s"), '\'', nil); len(got) != 0 {
		t.Fatal("'90s alone", got)
	}
}
