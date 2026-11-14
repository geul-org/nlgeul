//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingNames — the names of the rules a boundary is a place of joined with "/"; "-" when there is none (no rule on, or none fits)
package token

import "testing"

func TestSpacingNames(t *testing.T) {
	words := []step{{end: 3, word: spacingWord("아버지/열쇠:체언")}, {end: 5, word: spacingWord("가방/열쇠:체언")}}
	on := spacingScan("아버지가방", "2", "49", "50")
	if got := on.spacingNames(words, 0); got != "2/49/50" {
		t.Fatal(got)
	}
	off := spacingScan("아버지가방", "41")
	if got := off.spacingNames(words, 0); got != "-" {
		t.Fatal(got)
	}
}
