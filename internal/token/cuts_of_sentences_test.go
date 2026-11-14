//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCutsOfSentences — in everything a reader gives for the sentences the stream tests read, the branches of more than one word of a contest all end at the same End, and each branch cut by its End runs from the start of the stretch to its end without a gap, so that the places of the words fit the letters of the contest
package token

import "testing"

func TestCutsOfSentences(t *testing.T) {
	for _, words := range []string{"", "words-stream.tsv", "words.tsv"} {
		r := streamReader(t, words)
		if bad := eachSentence(r, func(r *Reader, s string) []string { return cutFaults(r.Contests(s)) }); len(bad) != 0 {
			t.Fatalf("%q: %q", words, bad)
		}
	}
}
