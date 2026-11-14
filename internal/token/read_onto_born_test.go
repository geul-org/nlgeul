//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReadOntoBorn — the letters after a sign of a pair are read by a scan of their own whose words are numbered on from where the whole unit's stopped: the particle carried over from the whole unit's way (no stem) has a number below that, the letters as a new word one at or above it
package token

import "testing"

func TestReadOntoBorn(t *testing.T) {
	ix, _ := newIndex(Options{})
	u := unit{text: ")를", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}}
	whole := unitScan(ix, u)
	whole.search()
	next := whole.registry().next
	got, _ := readOnto(ix, u, false)
	var moved, fresh Branch
	for _, b := range got[1].Branches {
		if b.Words[0].Attach {
			moved = b
		} else {
			fresh = b
		}
	}
	if moved.Words == nil || fresh.Words == nil || bornOf(moved) >= next || bornOf(fresh) < next {
		t.Fatal(next, bornOf(moved), bornOf(fresh))
	}
}
