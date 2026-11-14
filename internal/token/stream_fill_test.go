//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamFill — a stretch is read, one at a time, until a region is queued (a unit of several stretches gives them one by one); it does nothing when one is queued already, and at the end of the units reads no more
package token

import "testing"

func TestStreamFill(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	st := r.Stream("아버지 가방에. 들어가신다")
	st.fill()
	if begun(st) != 1 || len(st.queue) != 1 {
		t.Fatal(begun(st), len(st.queue))
	}
	st.fill()
	if begun(st) != 1 || len(st.queue) != 1 {
		t.Fatal("one is queued already")
	}
	st.queue = nil
	st.fill()
	if begun(st) != 2 || len(st.queue) != 1 {
		t.Fatal(begun(st), len(st.queue))
	}
	st.queue = nil
	st.fill()
	if begun(st) != 2 || len(st.queue) != 1 {
		t.Fatal("the second stretch of the unit", begun(st), len(st.queue))
	}
	for i := 0; i < 4 && (len(st.queue) > 0 || begun(st) < 3 || st.cur != nil); i++ {
		st.queue = nil
		st.fill()
	}
	st.queue = nil
	st.fill()
	if st.read != 3 || len(st.queue) != 0 || st.cur != nil {
		t.Fatal(st.read, len(st.queue))
	}
}
