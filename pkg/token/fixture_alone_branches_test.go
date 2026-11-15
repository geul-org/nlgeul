//ff:func feature=token type=helper control=iteration dimension=2
//ff:what aloneBranches — the candidates of the stretch that reads as the given text in the first way of a sentence read with the table words-alone.tsv, with the rules of the sentence or with none (ruled false), and every line of trace the stream wrote on the way
package token

import "testing"

func aloneBranches(t *testing.T, sentence, text string, ruled bool) ([]Branch, []string) {
	t.Helper()
	st := streamReader(t, "words-alone.tsv").StreamTrace(sentence)
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
