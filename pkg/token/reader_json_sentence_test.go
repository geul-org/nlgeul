//ff:func feature=token type=reader control=sequence
//ff:what TestReaderJSONSentence — a sentence set down as data: the stretches as Next sends them with every candidate of every word (nothing cut), places that match the letters without spaces, a stem no table knows has no name and is not known, a marker is known, the number given, and no stretch for a sentence of spaces only
package token

import "testing"

func TestReaderJSONSentence(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	s := "아버지가방에들어가신다 쪼간만"
	d := r.JSONSentence(s, 3)
	if d.Sentence != 3 || d.Text != s || len(d.Segments) != 2 || d.BestLambda == nil {
		t.Fatal(d)
	}
	first := d.Segments[0]
	if first.Rank != 1 || first.Ways == nil || len(first.Words) == 0 {
		t.Fatal(first)
	}
	if !jsonWordsFit(d, lettersOf(s)) {
		t.Fatal(d)
	}
	st := r.Stream(s)
	seg, _ := st.Next()
	if !jsonCandidatesAll(first, seg) {
		t.Fatal(first)
	}
	last := d.Segments[1].Words[0].Candidates[0].Parts
	if len(last) != 2 || last[0].Text != "쪼간" || last[0].Known || last[0].Name != nil || !last[1].Known || last[1].Name == nil || *last[1].Name != "보조사:만" {
		t.Fatal(last)
	}
	if e := r.JSONSentence("  ", 0); len(e.Segments) != 0 || e.Segments == nil {
		t.Fatal(e)
	}
}
