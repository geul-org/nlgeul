//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestNetEqualsOld — the stretches of every made-up sentence read by the net are those the old reading gives (a unit read whole, each contest cut by its branches): the same stretches, nodes with their candidates in order (letters, verdict, preference, number of birth), the words that follow each, how many ways, the ways one after another with their lambdas, and how many walks were skipped; with each reader (a table, spacing rules on and off, no table)
package parse

import "testing"

func TestNetEqualsOld(t *testing.T) {
	for name, r := range netReaders(t) {
		for _, s := range netSentences() {
			if d := netDiff(r, s); d != "" {
				t.Fatal(name, s, d)
			}
		}
	}
}
