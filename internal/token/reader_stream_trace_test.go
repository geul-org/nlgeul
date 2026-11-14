//ff:func feature=token type=reader control=sequence
//ff:what TestReaderStreamTrace — a stream made with StreamTrace keeps the trace lines of each unit it reads, the lines readUnit gives for it, and Notes hands them out when the unit has been read (one unit read for each Next), once; a stream made with Stream keeps none
package token

import (
	"slices"
	"testing"
)

func TestReaderStreamTrace(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	s := "아버지가 방에"
	units := unitsOf(r.ix, s, true)
	_, want0 := readUnit(r.ix, units[0], true)
	_, want1 := readUnit(r.ix, units[1], true)
	if len(want0) == 0 || len(want1) == 0 {
		t.Fatal("a trace has lines", len(want0), len(want1))
	}
	st := r.StreamTrace(s)
	if len(st.Notes()) != 0 {
		t.Fatal("nothing is read yet")
	}
	st.Next()
	if got := st.Notes(); !slices.Equal(got, want0) {
		t.Fatal(got, want0)
	}
	if len(st.Notes()) != 0 {
		t.Fatal("taken once")
	}
	st.Next()
	if got := st.Notes(); !slices.Equal(got, want1) {
		t.Fatal(got, want1)
	}
	plain := r.Stream(s)
	streamFirsts(plain)
	if len(plain.Notes()) != 0 {
		t.Fatal("no trace asked for")
	}
}
