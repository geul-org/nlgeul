//ff:func feature=token type=helper control=sequence
//ff:what complementStream — a stream over a sentence with the small table words-complement.tsv (되, 아니, 누구, 뭐 and a few nouns), its trace lines kept when asked for, and complementRule as the only rule of a sentence when ruled is true (none otherwise)
package token

import "testing"

func complementStream(t *testing.T, s string, trace, ruled bool) *Stream {
	t.Helper()
	r := streamReader(t, "words-complement.tsv")
	st := r.Stream(s)
	if trace {
		st = r.StreamTrace(s)
	}
	st.rules = nil
	if ruled {
		st.rules = []sentenceRule{complementRule()}
	}
	return st
}
