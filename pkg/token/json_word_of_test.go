//ff:func feature=token type=helper control=sequence
//ff:what TestJSONWordOf — the place, the letters and the glue of a node and a candidate for each branch of its contest
package token

import "testing"

func TestJSONWordOf(t *testing.T) {
	n := Node{Span: Span{Start: 2, End: 4}, Contest: Contest{Text: "물을", Glued: true, Branches: []Branch{{Verdict: 1}, {Verdict: 0.5}}}}
	w := jsonWordOf(n)
	if w.Span != [2]int{2, 4} || w.Text != "물을" || !w.Glued || len(w.Candidates) != 2 {
		t.Fatal(w)
	}
}
