//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestChainsEnded — a particle read only right after a closing (the 야 after a consonant) ends a chain only there: after the linking 다면 it is one more piece of the predicate's tail (가+(ㄴ)다면+야), and right on a stem that ends in a consonant it is not read (책야)
package parse

import "testing"

func TestChainsEnded(t *testing.T) {
	after := false
	for _, c := range scanOf("간다면야").chains(4, 0, -1) {
		last := c.pieces[len(c.pieces)-1]
		if last.names == "보조사:야" && len(c.pieces) < 2 {
			t.Fatal("야 right on a stem", c)
		}
		after = after || (last.names == "보조사:야" && c.close != nil)
	}
	if !after {
		t.Fatal("야 after 다면")
	}
	for _, c := range scanOf("책야").chains(2, 0, -1) {
		t.Fatal("a particle right on a stem that ends in a consonant", c)
	}
}
