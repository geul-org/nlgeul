//ff:func feature=token type=helper control=iteration dimension=2
//ff:what boundFirst — the first candidate of the word with the letters given in a sentence read with the small table words-bound.tsv and only the rules of a sentence given (none for the plain stream); "" when no word of the sentence has those letters
package token

import "testing"

func boundFirst(t *testing.T, sentence, word string, rules ...sentenceRule) string {
	t.Helper()
	st := streamReader(t, "words-bound.tsv").Stream(sentence)
	st.rules = rules
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		for _, n := range seg.Nodes {
			if n.Contest.Text == word {
				return n.Contest.Branches[0].String()
			}
		}
	}
	return ""
}
