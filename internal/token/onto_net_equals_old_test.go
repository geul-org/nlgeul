//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestOntoNetEqualsOld — the stretches of every made-up sentence with a sign of a pair with letters written onto it, read by the net, are those the old reading gives, but for the numbers of birth (which the net counts from another place, the words carried over from the whole unit first as before): the same stretches, nodes with their candidates in order, the words that follow each, how many ways, and the ways one after another; with each reader
package token

import "testing"

func TestOntoNetEqualsOld(t *testing.T) {
	for name, r := range netReaders(t) {
		for _, s := range ontoSentences() {
			if d := ontoDiff(r, s); d != "" {
				t.Fatal(name, s, d)
			}
		}
	}
}
