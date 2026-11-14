//ff:func feature=token type=reader control=sequence
//ff:what TestFitsNames — articles 49 and 50: a nominal with no part after it before a nominal fits; a nominal with a particle, a verb or a closing does not
package token

import "testing"

func TestFitsNames(t *testing.T) {
	if !fitsNames(gap{left: spacingWord("아버지/열쇠:체언"), right: spacingWord("가방/열쇠:체언", "에/부사격:에")}) || !fitsNames(gap{left: spacingWord("뚜루뚜"), right: spacingWord("솔")}) {
		t.Fatal("two nominals")
	}
	if fitsNames(gap{left: spacingWord("아버지/열쇠:체언", "가/주격:가"), right: spacingWord("가방/열쇠:체언")}) || fitsNames(gap{left: spacingWord("가/열쇠:동사"), right: spacingWord("가방/열쇠:체언")}) || fitsNames(gap{left: spacingWord("아버지/열쇠:체언"), right: spacingWord("가/열쇠:동사")}) {
		t.Fatal("a particle, a verb")
	}
}
