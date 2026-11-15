//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamFill — a stretch is read, one at a time, until a region is queued (a unit of several stretches gives them one by one); it does nothing when one is queued already, and at the end of the units reads no more. With a rule that waits for the predicate, a word with a waiting candidate makes it read on: the regions up to the predicate are queued together and judged (with no rule only the first is queued, and none is judged), and a sentence that ends with no predicate is read to its end and judged
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

	held := sentenceStream(t, "학생이 밥을 먹었다 학생이", false, waitingComplementRule())
	held.fill()
	if len(held.queue) != 3 || held.held || !held.queue[0].judged || !held.queue[2].judged || begun(held) != 3 {
		t.Fatal("read on to the predicate, and no further", len(held.queue), held.held, begun(held))
	}
	plain := sentenceStream(t, "학생이 밥을 먹었다", false)
	plain.fill()
	if len(plain.queue) != 1 || plain.queue[0].judged {
		t.Fatal("no rule, no wait", len(plain.queue))
	}
	open := sentenceStream(t, "학생이 밥을", false, waitingComplementRule())
	open.fill()
	if len(open.queue) != 2 || open.held || !open.queue[1].judged || open.read != 2 || open.cur != nil {
		t.Fatal("no predicate: read to the end", len(open.queue), open.held)
	}
}
