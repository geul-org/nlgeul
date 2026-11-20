//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what codeSpanFit — whether every stretch a stream sends for a sentence, and every way Again gives for it, has its region and its nodes in the places of the letters with all spaces left out: the stretches follow each other from 0 to the letters, the nodes fill the region, a node's letters without spaces are those at its place, and Again finds the stretch by the region sent
package parse

import "testing"

func codeSpanFit(t *testing.T, r *Reader, sentence string) {
	t.Helper()
	letters := lettersOf(sentence)
	st := r.Stream(sentence)
	end := 0
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		if seg.Region.Start != end || seg.Region.End <= seg.Region.Start {
			t.Fatal(sentence, "region", seg.Region, end)
		}
		end = seg.Region.End
		codeSpanAgain(t, st, sentence, seg, letters)
	}
	if end != len(letters) {
		t.Fatal(sentence, "end", end, len(letters))
	}
}
