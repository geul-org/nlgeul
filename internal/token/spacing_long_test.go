//ff:func feature=token type=reader control=sequence
//ff:what TestSpacingLong — a made-up unit of twenty letters full of cues (a new name, 겸, 씨, 님, 및, 등, 중) is read to the end with the rules on (the candidates 16 off, 408 on: 432 before a new word's stem was limited to newStemMax tokens)
package token

import "testing"

func TestSpacingLong(t *testing.T) {
	const in = "뚜루뚜겸갑돌씨와반장님및솔등박사중에"
	off, on := spacingReader(t, "words-spacing.tsv"), spacingReader(t, "words-spacing.tsv", SpacingArticles...)
	if got := len(off.Contests(in)[0].Branches); got != 16 {
		t.Fatal("off", got)
	}
	if got := len(on.Contests(in)[0].Branches); got != 408 {
		t.Fatal("on", got)
	}
}
