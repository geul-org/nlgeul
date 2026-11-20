//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what ruledBranches — the candidates of the first stretch that reads as the given text in the first way of a sentence read with the given table of testdata, with the rules of the sentence or with none (ruled false), and every line of trace the stream wrote on the way
package parse

import "testing"

func ruledBranches(t *testing.T, table, sentence, text string, ruled bool) ([]Branch, []string) {
	t.Helper()
	st := streamReader(t, table).StreamTrace(sentence)
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
