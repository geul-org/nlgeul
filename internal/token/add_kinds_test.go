//ff:func feature=token type=rule control=sequence
//ff:what TestAddKinds — a graph for each kind, each with its exception and two rebuttals
package token

import "testing"

func TestAddKinds(t *testing.T) {
	ix := &index{}
	ix.addKinds()
	if ix.graphs != 6 || ix.edges != 18 || len(ix.kinds) != 6 || ix.seen == nil {
		t.Fatal(ix.graphs, ix.edges)
	}
}
