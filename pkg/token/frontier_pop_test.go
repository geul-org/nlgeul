//ff:func feature=token type=helper control=sequence
//ff:what TestFrontierPop — the first walk comes out and the others stay in order
package token

import "testing"

func TestFrontierPop(t *testing.T) {
	f := frontier{items: []*walk{{node: 3}, {node: 4}}}
	if w := f.pop(); w.node != 3 || len(f.items) != 1 || f.items[0].node != 4 {
		t.Fatal(w, f.items)
	}
}
