//ff:func feature=token type=reader control=sequence
//ff:what TestRootGrown — the root is grown once by each way the stem from 0 to n stands (a word of two kinds has two readings, a new word one), each a path from the root with that word as its last
package token

import "testing"

func TestRootGrown(t *testing.T) {
	root := &path{verdict: 1}
	sc := wordsScan("지금", layerCost)
	sc.ix.lex["지금"] = word{kinds: kindNominal | kindAdverb, count: 100, by: [6]int{60, 0, 0, 40}}
	got := sc.rootGrown(root, 2, tail{start: 2})
	if len(got) != 2 || got[0].prev != root || got[1].prev != root || got[0].last.sig != "지금[열쇠:체언]" || got[1].last.sig != "지금[열쇠:부사]" {
		t.Fatal(got)
	}
	one := scanOf("고양이").rootGrown(root, 3, tail{start: 3})
	if len(one) != 1 || one[0].prev != root || one[0].sig != "고양이" || one[0].last.end != 3 {
		t.Fatal("a new word is one reading", one)
	}
}
