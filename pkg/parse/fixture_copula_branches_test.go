//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what copulaBranches — the candidates of the stretch that reads as the given text in the first way of a sentence read with the table words-copula.tsv, with the rules of the sentence or with none (ruled false), and every line of trace the stream wrote on the way
package parse

import "testing"

func copulaBranches(t *testing.T, sentence, text string, ruled bool) ([]Branch, []string) {
	t.Helper()
	st := streamReader(t, "words-copula.tsv").StreamTrace(sentence)
	if !ruled {
		st.rules = nil
	}
	var found []Branch
	var notes []string
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		notes = append(notes, st.Notes()...)
		for _, n := range seg.Nodes {
			if n.Contest.Text == text && found == nil {
				found = n.Contest.Branches
			}
		}
	}
	return found, notes
}
