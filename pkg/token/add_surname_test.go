//ff:func feature=token type=rule control=sequence
//ff:what TestAddSurname — the family names' graph is one graph with no rebuttal: its verdict is 1 and the index counts one graph and no edge more
package token

import "testing"

func TestAddSurname(t *testing.T) {
	ix := &index{}
	ix.addSurname()
	if ix.surname != 1 || ix.graphs != 1 || ix.edges != 0 {
		t.Fatal(ix.surname, ix.graphs, ix.edges)
	}
}
