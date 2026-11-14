//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReadOntoNetTrace — the trace of a sign of a pair with letters written onto it read by the net has the head's lines once, and lines of the letters read; the whole unit is read only for the words that begin at its head or after the sign, so it has fewer lines than the old reading, whose lines of the letters after the sign it has all
package token

import "testing"

func TestReadOntoNetTrace(t *testing.T) {
	ix, _ := newIndex(Options{})
	u := unit{text: ")를", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}, notes: []string{"head"}}
	_, want := readOnto(ix, u, true)
	_, got := readOntoNet(ix, u, true)
	count := 0
	for _, n := range got {
		if n == "head" {
			count++
		}
	}
	if count != 1 || len(got) < 2 || len(got) > len(want) {
		t.Fatal(count, len(got), len(want), got)
	}
}
