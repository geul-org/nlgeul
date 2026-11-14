//ff:func feature=token type=helper control=sequence
//ff:what TestNewRegistry — a registry begins with no words, the number given as the next to give, empty tables and a root that stands (verdict 1) with nothing before it
package token

import "testing"

func TestNewRegistry(t *testing.T) {
	r := newRegistry(7)
	if r.next != 7 || len(r.born) != 0 || len(r.states) != 0 || len(r.all) != 0 || len(r.lone) != 0 || r.twins != 0 {
		t.Fatal(r)
	}
	if r.root == nil || r.root.verdict != 1 || r.root.prev != nil {
		t.Fatal("the root", r.root)
	}
}
