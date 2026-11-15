//ff:func feature=token type=helper control=sequence
//ff:what sentenceStream — a stream over a sentence with the small table words-sentence.tsv (학생, 되, 먹, 밥), its trace lines kept when asked for, and the rules of a sentence given in place of the index's (none when none is given), so that a test does not hang on what the index holds
package token

import "testing"

func sentenceStream(t *testing.T, s string, trace bool, rules ...sentenceRule) *Stream {
	t.Helper()
	r := streamReader(t, "words-sentence.tsv")
	st := r.Stream(s)
	if trace {
		st = r.StreamTrace(s)
	}
	st.rules = rules
	return st
}
