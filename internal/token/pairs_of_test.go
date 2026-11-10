//ff:func feature=token type=reader control=sequence
//ff:what TestPairsOf — brackets pair by nesting and quotation marks two by two, across spaces; a bracket and a quotation mark inside a span of code are no signs; a sign with no partner is left out; a quote closed in a later word with a particle after it pairs, an apostrophe between Latin letters (it's) and a lone '90s do not
package token

import "testing"

func TestPairsOf(t *testing.T) {
	got := pairsOf([]rune("(가 \"나\") `(\"` 다)"))
	if len(got) != 6 || got[0] != 6 || got[6] != 0 || got[3] != 5 || got[8] != 11 || got[11] != 8 {
		t.Fatal(got)
	}
	if _, ok := got[9]; ok {
		t.Fatal("a bracket inside code")
	}
	if got = pairsOf([]rune("곰이 '사자 호랑이'는 싫다")); len(got) != 2 || got[3] != 10 || got[10] != 3 {
		t.Fatal("a closing quote with a particle after it", got)
	}
	if got = pairsOf([]rune("it's '90s")); len(got) != 0 {
		t.Fatal("it's, '90s", got)
	}
}
