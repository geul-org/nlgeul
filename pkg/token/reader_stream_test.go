//ff:func feature=token type=reader control=sequence
//ff:what TestReaderStream — a stream over words written apart: Next gives the stretches in order (the units of Reader.Contests, as many), each tiles its region from node to node and holds one word in every candidate, the node that is the whole region holds the candidates of one word of the contest with their letters, verdict and preference, the ways of the units differ in number
package token

import "testing"

func TestReaderStream(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	s := "아버지 가방에 들어가신다"
	letters := []rune("아버지가방에들어가신다")
	st := r.Stream(s)
	segs := streamFirsts(st)
	if len(segs) != 3 || segs[0].Region != (Span{0, 3}) || segs[1].Region != (Span{3, 6}) || segs[2].Region != (Span{6, 11}) {
		t.Fatal(segs)
	}
	if segs[0].Ways.Int64() != 1 || segs[1].Ways.Int64() != 1 || segs[2].Ways.Int64() != 8 {
		t.Fatal(segs[0].Ways, segs[1].Ways, segs[2].Ways)
	}
	if bad := waysCheck(segs, letters); len(bad) != 0 {
		t.Fatal(bad)
	}
	if bad := regionMatches(st, r.Contests(s)); len(bad) != 0 {
		t.Fatal(bad)
	}
	if segmentLine(segs[2]) != "들어가신다[6-11]" || segs[2].Rank != 1 {
		t.Fatal(segmentLine(segs[2]))
	}
	if _, ok := st.Next(); ok {
		t.Fatal("the end stays the end")
	}
}
