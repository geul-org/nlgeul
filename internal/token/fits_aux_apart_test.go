//ff:func feature=token type=reader control=sequence
//ff:what TestFitsAuxApart — the proviso of article 47: an auxiliary of the lists after a closing and a particle fits; so do the stems of a word with a particle and an auxiliary together made of a dependent noun; no particle after the closing does not
package token

import "testing"

func TestFitsAuxApart(t *testing.T) {
	if !fitsAuxApart(gap{left: spacingWord("읽/열쇠:동사", "어/맺음:어", "도/보조사:도"), right: spacingWord("보/열쇠:동사")}) {
		t.Fatal("읽어도 보")
	}
	if !fitsAuxApart(gap{left: spacingWord("듯/열쇠:체언", "도/보조사:도"), right: spacingWord("하/열쇠:동사")}) {
		t.Fatal("듯도 하")
	}
	if fitsAuxApart(gap{left: spacingWord("읽/열쇠:동사", "어/맺음:어"), right: spacingWord("보/열쇠:동사")}) || fitsAuxApart(gap{left: spacingWord("듯/열쇠:체언"), right: spacingWord("하/열쇠:동사")}) || fitsAuxApart(gap{left: spacingWord("뚜루뚜", "도/보조사:도"), right: spacingWord("하/열쇠:동사")}) {
		t.Fatal("no particle, no auxiliary")
	}
}
