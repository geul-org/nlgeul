//ff:func feature=token type=reader control=sequence
//ff:what TestMarkGap — the boundary after a word of a candidate: the two words that meet and the letters each takes up, the first word counted from the unit's head
package token

import "testing"

func TestMarkGap(t *testing.T) {
	a, b, c := spacingWord("아버지/열쇠:체언"), spacingWord("가방/열쇠:체언"), spacingWord("들어가/열쇠:동사")
	words := []step{{end: 3, word: a}, {end: 5, word: b}, {end: 8, word: c}}
	if g := markGap(words, 0); g.leftLen != 3 || g.rightLen != 2 || g.left.Parts[0].Text != "아버지" || g.right.Parts[0].Text != "가방" {
		t.Fatal(g)
	}
	if g := markGap(words, 1); g.leftLen != 2 || g.rightLen != 3 || g.right.Parts[0].Text != "들어가" {
		t.Fatal(g)
	}
}
