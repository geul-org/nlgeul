//ff:func feature=parse type=rule control=sequence
//ff:what TestUnitEntries — the letters of a unit of the unit table give an entry named as the table names the unit (㎡: 단위:제곱미터), the same entry when asked again, with a graph; letters that are no unit give none
package parse

import "testing"

func TestUnitEntries(t *testing.T) {
	ix := &index{formEntries: map[string]*entry{}}
	got := ix.unitEntries("㎡")
	if len(got) != 1 || got[0].name != "단위:제곱미터" || ix.graphs != 1 || got[0].g == nil {
		t.Fatal(got, ix.graphs)
	}
	if again := ix.unitEntries("㎡"); again[0] != got[0] || ix.graphs != 1 {
		t.Fatal("one entry for one name")
	}
	if none := ix.unitEntries("ⓐ"); none != nil {
		t.Fatal(none)
	}
}
