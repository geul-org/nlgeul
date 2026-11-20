//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what codeSpanAgain — checks a segment and the ways Again gives after it for its region (up to three) with codeSpanNodes, each way keeping the region that was sent
package parse

import "testing"

func codeSpanAgain(t *testing.T, st *Stream, sentence string, seg Segment, letters []rune) {
	t.Helper()
	for tries := 0; tries < 3; tries++ {
		codeSpanNodes(t, sentence, seg, letters)
		next, more := st.Again(seg.Region)
		if !more {
			return
		}
		if next.Region != seg.Region {
			t.Fatal(sentence, "again", next.Region, seg.Region)
		}
		seg = next
	}
}
