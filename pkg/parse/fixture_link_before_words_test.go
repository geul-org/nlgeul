//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what linkBeforeWords — the candidates of every word of a sentence read with words-link-before-aux.tsv, in the order of the words, with the rules of the index or none (rules false), and the trace lines Notes gave after each Next
package parse

import "testing"

func linkBeforeWords(t *testing.T, sentence string, rules bool) ([][]Branch, []string) {
	t.Helper()
	r := streamReader(t, "words-link-before-aux.tsv")
	st := r.StreamTrace(sentence)
	if !rules {
		st.rules = nil
	}
	var words [][]Branch
	var notes []string
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		for _, n := range seg.Nodes {
			words = append(words, n.Contest.Branches)
		}
		notes = append(notes, st.Notes()...)
	}
	return words, notes
}
