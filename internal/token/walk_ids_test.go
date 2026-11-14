//ff:func feature=token type=helper control=sequence
//ff:what TestWalkIds — the places of the nodes of a walk, first to last
package token

import "testing"

func TestWalkIds(t *testing.T) {
	w := &walk{prev: &walk{prev: &walk{node: 4}, node: 7}, node: 2}
	if got := w.ids(); len(got) != 3 || got[0] != 4 || got[1] != 7 || got[2] != 2 {
		t.Fatal(got)
	}
}
