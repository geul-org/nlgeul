//ff:func feature=token type=reader control=sequence
//ff:what TestStreamAgain — 아버지가방에들어가신다 is one stretch with no agreed place inside, so a stretch of the whole letters, [0-11] (the stream does not cut it at 6); its ways come one after the other through Again, each better or as good as the one before, all of the ways of cutting its contest makes, once each, and then none; a stretch that is not one Next has sent (not yet, or not at all) gives nothing; Next after them gives the end; with a closing dot Next gives the dot's stretch
package token

import (
	"math/big"
	"testing"
)

func TestStreamAgain(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	s := "아버지가방에들어가신다"
	st := r.Stream(s)
	first, ok := st.Next()
	if !ok || first.Region != (Span{0, 11}) || first.Rank != 1 || first.Ways.Cmp(big.NewInt(60)) != 0 || segmentLine(first) != "아버지[0-3] 가방에[3-6] 들어가신다[6-11]" {
		t.Fatal(first, ok, segmentLine(first))
	}
	ways := streamWays(st, first)
	contests := r.Contests(s)
	forms := formsOf(cutsOf(contests[0], 0))
	keys := segmentKeys(ways)
	if len(ways) != 60 || len(forms) != 60 || len(keys) != 60 || !ranked(ways) {
		t.Fatal(len(ways), len(forms), len(keys))
	}
	if bad := missingForms(forms, keys); len(bad) != 0 {
		t.Fatal(bad)
	}
	if _, ok := st.Again(first.Region); ok {
		t.Fatal("none is left, and it stays so")
	}
	if _, ok := st.Again(Span{0, 3}); ok {
		t.Fatal("not a stretch of the stream")
	}
	if _, ok := st.Next(); ok {
		t.Fatal("the end")
	}

	dotted := r.Stream(s + ".")
	d, _ := dotted.Next()
	streamWays(dotted, d)
	if _, ok := dotted.Again(Span{11, 12}); ok {
		t.Fatal("not sent yet")
	}
	dot, ok := dotted.Next()
	if !ok || dot.Region != (Span{11, 12}) || dot.Ways.Cmp(big.NewInt(1)) != 0 || len(dot.Nodes) != 1 {
		t.Fatal(dot, ok)
	}
	if _, ok := dotted.Again(Span{11, 12}); ok {
		t.Fatal("one way only")
	}

	cut := r.Stream(redoText)
	cut.Next()
	cut.Next()
	one, _ := cut.Next()
	two, ok := cut.Again(one.Region)
	if want := redoWays(t, r, 3); !ok || segmentLine(two) != want[1] || two.Rank != 2 {
		t.Fatal("a stretch a window cut off is read again", segmentLine(two), want)
	}
}
