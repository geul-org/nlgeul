//ff:func feature=token type=rule control=sequence
//ff:what TestNewIndexTypo — the index has a graph for each closing written wrong (17 more than the closings alone: 486 without a word table, 535 with one), the edges that go with them (1014 and 1155), and the total of counts and the set of closings are those of the right closings
package token

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
	if bare.graphs != 486 || bare.edges != 1014 || words.graphs != 535 || words.edges != 1155 {
		t.Fatal(bare.graphs, bare.edges, words.graphs, words.edges)
	}
	if bare.closingSet["께"] || bare.closingSet["ㄹ께"] || bare.closingSet["읍니다"] || bare.closingSet["ㄹ게"] != true {
		t.Fatal("the set of closings is those of the tables")
	}
}
