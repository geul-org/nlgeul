//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddSigns — each named sign gets a graph for each of its readings in the order they are printed; the full stop's graphs are the closing readings of the stops; 16 readings in all
package token

import (
	"reflect"
	"testing"
)

func TestAddSigns(t *testing.T) {
	ix := &index{}
	ix.addDots()
	graphs := ix.graphs
	ix.addSigns()
	if ix.graphs-graphs != 16 || len(ix.marks) != len(signOrderNamed)+1 {
		t.Fatal(ix.graphs-graphs, len(ix.marks))
	}
	if len(ix.marks['.']) != 2 || ix.marks['.'][0].name != "마침표" || ix.marks['.'][1].name != "줄임표" {
		t.Fatal("the full stop", ix.marks['.'])
	}
	for sign, want := range map[rune][]string{',': {"쉼표"}, '#': {"제목", "샵", "태그"}, '@': {"골뱅이", "명칭"}, '~': {"물결표"}} {
		if got := entryNames(ix.marks[sign]); !reflect.DeepEqual(got, want) {
			t.Errorf("%c: %q, want %q", sign, got, want)
		}
	}
}
