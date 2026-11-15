//ff:func feature=token type=helper control=iteration dimension=2
//ff:what oneCandidates — the candidates of the first word 한 of a sentence read with words-one.tsv (the best way of each stretch), with the rules of the index or none (rules false) and with a trace or not, and every trace line Notes gave (the test stops when no word is 한)
package token

import "testing"

func oneCandidates(t *testing.T, sentence string, rules, trace bool) ([]Branch, []string) {
	t.Helper()
	r := streamReader(t, "words-one.tsv")
	st := r.Stream(sentence)
	if trace {
		st = r.StreamTrace(sentence)
	}
	if !rules {
		st.rules = nil
	}
	var brs []Branch
	var notes []string
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		notes = append(notes, st.Notes()...)
		for _, n := range seg.Nodes {
			if brs == nil && n.Contest.Text == oneStem {
				brs = n.Contest.Branches
			}
		}
	}
	if brs == nil {
		t.Fatalf("no word 한 in %q", sentence)
	}
	return brs, notes
}
