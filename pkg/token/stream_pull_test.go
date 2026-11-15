//ff:func feature=token type=reader control=sequence
//ff:what TestStreamPull — the next unit (all its stretches pulled, pullUnit) is read into regions in order: each contest of it at its place counted without spaces, marked with its unit and Glued as the old reading marks them (unitContests) (a unit after a space is not glued, a sign written right on it is), the cursor after its letters, the space after it noted for the next unit; the trace lines made in reading a stretch go with its region and none is left pending, and a region with a candidate that waits for the predicate makes the stream wait (hold)
package token

import "testing"

func TestStreamPull(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	st := r.Stream("아버지. 방에들어(가신다)")
	pullUnit(st)
	if len(st.queue) != 2 || st.queue[0].span != (Span{0, 3}) || st.queue[1].span != (Span{3, 4}) || st.cursor != 4 || st.read != 1 || !st.spaced {
		t.Fatal(st.queue, st.cursor)
	}
	if a, b := st.queue[0].nodes[0].Contest, st.queue[1].nodes[0].Contest; a.Unit != 0 || a.Glued || b.Unit != 0 || !b.Glued {
		t.Fatal(a, b)
	}
	pullUnit(st)
	if len(st.queue) != 3 || st.queue[2].span != (Span{4, 8}) {
		t.Fatal(len(st.queue))
	}
	if first := st.queue[2].nodes[0].Contest; first.Unit != 1 || first.Glued {
		t.Fatal(first)
	}
	pullUnit(st)
	if c := st.queue[len(st.queue)-1].nodes[0].Contest; c.Unit != 2 || !c.Glued {
		t.Fatal("a unit right after another with no space is glued", c)
	}

	far := streamReader(t, "words-far.tsv").Stream(windowText(5))
	far.pull()
	far.pull()
	if len(far.queue) != 2 || far.cur == nil || far.queue[1].redo == nil || far.queue[1].redo.at != 4 || far.queue[1].redo.head.Text != "마당에서" || far.queue[1].redo.head.Glued != true || far.cursor != 8 {
		t.Fatal("a stretch of a unit is queued as it is read, with what is needed to read it again", far.queue)
	}

	traced := sentenceStream(t, "학생이 되었다", true, complementRule())
	traced.pull()
	if len(traced.queue) != 1 || len(traced.queue[0].notes) == 0 || traced.pending != nil || !traced.held {
		t.Fatal("the lines of the stretch go with its region", traced.queue[0].notes, traced.pending, traced.held)
	}
	traced.pull()
	traced.pull()
	if len(traced.queue) != 2 || len(traced.queue[1].notes) == 0 || traced.pending != nil || traced.held {
		t.Fatal("the predicate ends the wait", traced.held)
	}
}
