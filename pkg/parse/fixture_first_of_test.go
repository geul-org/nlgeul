//ff:func feature=parse type=helper control=sequence
//ff:what firstOf — the first word of the best way of a region
package parse

import "testing"

func firstOf(t *testing.T, rg *region) Node {
	t.Helper()
	seg, ok := rg.peek()
	if !ok || len(seg.Nodes) == 0 {
		t.Fatal("no way", rg.span)
	}
	return seg.Nodes[0]
}
