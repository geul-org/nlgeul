//ff:func feature=parse type=rule control=sequence
//ff:what TestAddSwap — the graph of a round quotation mark read against its shape is put under the sign after its other graphs, as an opening (열기) for a mark drawn as a closing one and a closing (닫기) for a mark drawn as an opening one, with the sign after a colon in its name, and the graphs and edges are counted
package parse

import (
	"reflect"
	"testing"
)

func TestAddSwap(t *testing.T) {
	ix := &index{pairs: map[rune][]*entry{}}
	ix.addPair('”', pairRows['”'], shutName)
	ix.addSwap('”')
	ix.addSwap('‘')
	if got := entryNames(ix.pairs['”']); !reflect.DeepEqual(got, []string{"닫기:”", "열기:”"}) {
		t.Fatal(got)
	}
	swapped := ix.pairs['‘']
	if len(swapped) != 1 || swapped[0].name != "닫기:‘" || swapped[0].form != shutName || swapped[0].g == nil || ix.pairs['”'][1].form != openName {
		t.Fatal(swapped)
	}
	if ix.graphs != 3 || ix.edges != 2+1+1 {
		t.Fatal(ix.graphs, ix.edges)
	}
}
