//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamRedoAgain — the ways of a stretch a window cut off come one after another as the stretch itself gave them; going to another such stretch and back passes over the ways already sent; once the ways are used up none is given, and only the one stretch read again last is held
package token

import "testing"

func TestStreamRedoAgain(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	want := redoWays(t, r, 5)
	st := r.Stream(redoText)
	st.Next()
	st.Next()
	a, _ := st.Next()
	st.Next()
	st.Next()
	b, _ := st.Next()
	for k := 1; k < 3; k++ {
		if seg, ok := st.Again(a.Region); !ok || segmentLine(seg) != want[k] {
			t.Fatal("way", k, segmentLine(seg), want[k])
		}
	}
	held := st.redone
	if _, ok := st.Again(b.Region); !ok || st.redone == held || st.redoOf == nil {
		t.Fatal("another stretch is read again")
	}
	if seg, ok := st.Again(a.Region); !ok || segmentLine(seg) != want[3] || seg.Rank != 4 {
		t.Fatal("back to the first, past the ways sent", segmentLine(seg))
	}
	for ok := true; ok; _, ok = st.Again(a.Region) {
	}
	if _, ok := st.Again(a.Region); ok {
		t.Fatal("used up")
	}
}
