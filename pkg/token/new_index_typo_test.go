//ff:func feature=token type=rule control=sequence
//ff:what TestNewIndexTypo — the index has a graph for each closing written wrong (17 more than the closings alone: 652 without a word table, 708 with one, which also reads two forms of 거 written with 꺼), the edges that go with them (1551 and 1709), and the total of counts and the set of closings are those of the right closings
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
	if bare.graphs != 652 || bare.edges != 1551 || words.graphs != 708 || words.edges != 1709 {
		t.Fatal(bare.graphs, bare.edges, words.graphs, words.edges)
	}
	if bare.closingSet["께"] || bare.closingSet["ㄹ께"] || bare.closingSet["읍니다"] || bare.closingSet["ㄹ게"] != true {
		t.Fatal("the set of closings is those of the tables")
	}
}
