//ff:func feature=parse type=reader control=sequence
//ff:what TestIndexPoliteNamed — with a word table the names of the table's polite slots are the index's; another slot's name is not, and without a word table none is
package parse

import "testing"

func TestIndexPoliteNamed(t *testing.T) {
	ix := wordsScan("가", layerCost).ix
	if !ix.politeNamed("선어말:오") || !ix.politeNamed("선어말:자옵") || ix.politeNamed("선어말:시") || len(ix.polite) != len(politeSlots) {
		t.Fatal(len(ix.polite))
	}
	if scanOf("가").ix.politeNamed("선어말:오") {
		t.Fatal("no word table")
	}
}
