//ff:func feature=token type=reader control=sequence
//ff:what TestStreamPull — the next unit (all its stretches pulled, pullUnit) is read into regions in order: each contest of it at its place counted without spaces, marked with its unit and Glued as Contests marks them (a unit after a space is not glued, a sign written right on it is), the cursor after its letters, the space after it noted for the next unit
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
}
