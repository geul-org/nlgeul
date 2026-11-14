//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestStreamAgainSentences — for the sentences the stream tests read, with the small tables and with none, every stretch gives all its ways through Again: the same ways of cutting as the branches of its contest make, each once, in ranks 1, 2, 3 … with lambdas that never rise, each tiling its stretch with candidates of one word, and the end of the sentence after the last stretch
package token

import (
	"strings"
	"testing"
)

func TestStreamAgainSentences(t *testing.T) {
	for _, words := range []string{"", "words-stream.tsv", "words.tsv"} {
		r := streamReader(t, words)
		check := func(r *Reader, s string) []string {
			return waysMatch(r.Stream(s), r.Contests(s), []rune(strings.ReplaceAll(s, " ", "")))
		}
		if bad := eachSentence(r, check); len(bad) != 0 {
			t.Fatalf("%q: %q", words, bad)
		}
	}
}
