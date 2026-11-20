//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestStretchScores — in everything the old reading (unitContests) gives for the sentences the stream tests read, with the small tables and without any, each branch of more than one word holds the verdict and the preference of each of its words, the least verdict and the sum of the preferences being the branch's own
package parse

import "testing"

func TestStretchScores(t *testing.T) {
	for _, words := range []string{"", "words-stream.tsv", "words.tsv"} {
		r := streamReader(t, words)
		if bad := eachSentence(r, func(r *Reader, s string) []string { return scoreFaults(unitContests(r, s)) }); len(bad) != 0 {
			t.Fatalf("%q: %q", words, bad)
		}
	}
}
