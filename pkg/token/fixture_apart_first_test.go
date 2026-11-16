//ff:func feature=token type=helper control=iteration dimension=2
//ff:what apartFirst — the first candidate of the word with the letters given in a sentence read with the small table words-apart.tsv, a trace and only the rules of a sentence given (none for the plain stream), with every trace line Notes gave; "" when no word of the sentence has those letters
package token

import "testing"

func apartFirst(t *testing.T, sentence, word string, rules ...sentenceRule) (string, []string) {
	t.Helper()
	st := streamReader(t, "words-apart.tsv").StreamTrace(sentence)
	st.rules = rules
	first := ""
	var notes []string
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		notes = append(notes, st.Notes()...)
		for _, n := range seg.Nodes {
			if first == "" && n.Contest.Text == word {
				first = n.Contest.Branches[0].String()
			}
		}
	}
	return first, notes
}
