//ff:func feature=token type=reader control=sequence
//ff:what TestStreamAgainHead — AgainHead gives the next way of a stretch with the first head candidates of each node, the same way Again gives; Again after it goes on with the way after (the head does not move or change the stream); a stretch that is not one Next has sent gives nothing; the stretches a window cut off, read again for AgainHead, give the first head of what Again gives
package token

import "testing"

func TestStreamAgainHead(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	s := "아버지가방에들어가신다"
	st, full := r.Stream(s), r.Stream(s)
	first, _ := st.NextHead(1)
	f1, _ := full.Next()
	second, ok := st.AgainHead(first.Region, 1)
	f2, _ := full.Again(f1.Region)
	if !ok || second.Rank != 2 || len(headMatch(f2, second, 1)) != 0 {
		t.Fatal(second, ok, headMatch(f2, second, 1))
	}
	third, _ := st.Again(first.Region)
	f3, _ := full.Again(f1.Region)
	if third.Rank != 3 || len(headMatch(f3, third, 0)) != 0 {
		t.Fatal("Again after AgainHead", third.Rank)
	}
	if _, ok := st.AgainHead(Span{0, 3}, 1); ok {
		t.Fatal("not a stretch of the stream")
	}
	if bad := headWalk(r.Stream(redoText), r.Stream(redoText), 1); len(bad) != 0 {
		t.Fatal(bad)
	}
}
