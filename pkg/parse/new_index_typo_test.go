//ff:func feature=parse type=rule control=sequence
//ff:what TestNewIndexTypo — the index has a graph for each marker of the grammar tables, for each closing written wrong, for the family names (addSurname) and, with a word table, for the forms that need one; the numbers of graphs and of edges are pinned here, without and with a word table, so that a row added to a table or a graph added to the index shows as a change of these numbers; the set of closings is those of the tables (a form written wrong is not in it)
package parse

import "testing"

func TestNewIndexTypo(t *testing.T) {
	bare, err := newIndex(Options{})
	if err != nil {
		t.Fatal(err)
	}
	words, err := newIndex(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	if bare.graphs != 891 || bare.edges != 2438 || words.graphs != 1005 || words.edges != 2813 {
		t.Fatal(bare.graphs, bare.edges, words.graphs, words.edges)
	}
	if bare.closingSet["께"] || bare.closingSet["ㄹ께"] || bare.closingSet["읍니다"] || bare.closingSet["ㄹ게"] != true {
		t.Fatal("the set of closings is those of the tables")
	}
}
