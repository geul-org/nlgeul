//ff:func feature=token type=reader control=sequence
//ff:what TestFitsAux — article 47: an auxiliary of the lists after a word that ends in a closing fits, and an auxiliary made of a dependent noun after a word that ends in a modifying closing; neither after the other kind of closing, nor a stem that is no auxiliary
package token

import "testing"

func TestFitsAux(t *testing.T) {
	closing, modifying := spacingWord("읽/열쇠:동사", "어/맺음:어"), spacingWord("올/열쇠:동사", "ㄹ/꾸밈:ㄹ")
	if !fitsAux(gap{left: closing, right: spacingWord("보/열쇠:동사", "았/선어말:았")}) || !fitsAux(gap{left: closing, right: spacingWord("싶/열쇠:형용사")}) {
		t.Fatal("after a closing")
	}
	if !fitsAux(gap{left: modifying, right: spacingWord("듯하/열쇠:형용사")}) {
		t.Fatal("after a modifying closing")
	}
	if fitsAux(gap{left: modifying, right: spacingWord("보/열쇠:동사")}) || fitsAux(gap{left: closing, right: spacingWord("듯하/열쇠:형용사")}) || fitsAux(gap{left: closing, right: spacingWord("사과/열쇠:체언")}) {
		t.Fatal("the other kind of closing, no auxiliary")
	}
}
