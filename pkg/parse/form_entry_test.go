//ff:func feature=parse type=rule control=sequence
//ff:what TestFormEntry — a name gets one entry with that name and a graph, counted once; the same name gives the same entry and another name another
package parse

import "testing"

func TestFormEntry(t *testing.T) {
	ix := &index{formEntries: map[string]*entry{}}
	a := ix.formEntry("원:로마자")
	if a.name != "원:로마자" || a.g == nil || ix.graphs != 1 {
		t.Fatal(a, ix.graphs)
	}
	if ix.formEntry("원:로마자") != a || ix.graphs != 1 {
		t.Fatal("one entry for one name")
	}
	if b := ix.formEntry("원:한글"); b == a || ix.graphs != 2 {
		t.Fatal(b, ix.graphs)
	}
}
