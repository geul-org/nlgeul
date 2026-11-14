//ff:func feature=token type=reader control=sequence
//ff:what TestStreamFill — units are read one at a time until a region is queued (a unit of several contests queues them all); it does nothing when one is queued already, and at the end of the units reads no more
package token

import "testing"

func TestStreamFill(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	st := r.Stream("아버지 가방에. 들어가신다")
	st.fill()
	if st.read != 1 || len(st.queue) != 1 {
		t.Fatal(st.read, len(st.queue))
	}
	st.fill()
	if st.read != 1 || len(st.queue) != 1 {
		t.Fatal("one is queued already")
	}
	st.queue = nil
	st.fill()
	if st.read != 2 || len(st.queue) != 2 {
		t.Fatal(st.read, len(st.queue))
	}
	st.queue = nil
	st.fill()
	st.queue = nil
	st.fill()
	if st.read != 3 || len(st.queue) != 0 {
		t.Fatal(st.read, len(st.queue))
	}
}
