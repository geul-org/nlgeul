//ff:func feature=token type=helper control=sequence
//ff:what TestTwos — the places of a sign paired two by two, the odd last one left out; a closing one with a particle written onto it pairs with the one opened a word before
package token

import "testing"

func TestTwos(t *testing.T) {
	got := twos([]rune("`가` `나"), '`')
	if len(got) != 2 || got[0] != 2 || got[2] != 0 {
		t.Fatal(got)
	}
	if got = twos([]rune("곰이 `사자 호랑이`는 싫다"), '`'); len(got) != 2 || got[3] != 10 || got[10] != 3 {
		t.Fatal("a closing one with a particle after it", got)
	}
	if got = twos([]rune("`가"), '`'); len(got) != 0 {
		t.Fatal("one alone", got)
	}
}
