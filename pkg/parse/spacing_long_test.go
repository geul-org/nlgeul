//ff:func feature=parse type=reader control=sequence
//ff:what TestSpacingLong — a made-up unit of twenty letters full of cues (a new name, 겸, 씨, 님, 및, 등, 중) is read to the end with the rules on (the candidates, the ways through the net the stream reads it by (netLines), 20 off, 568 on: 16 and 408 before 등 was read as a unit with no number written onto it as well, 432 before a new word's stem was limited to newStemMax tokens)
package parse

import "testing"

func TestSpacingLong(t *testing.T) {
	const in = "뚜루뚜겸갑돌씨와반장님및솔등박사중에"
	off, on := spacingReader(t, "words-spacing.tsv"), spacingReader(t, "words-spacing.tsv", SpacingArticles...)
	if got := len(netLines(off, in)); got != 20 {
		t.Fatal("off", got)
	}
	if got := len(netLines(on, in)); got != 568 {
		t.Fatal("on", got)
	}
}
