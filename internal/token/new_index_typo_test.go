//ff:func feature=token type=rule control=sequence
//ff:what TestNewIndexTypo — the index has a graph for each closing written wrong (17 more than the closings alone: 521 without a word table, 572 with one, which also reads two forms of 거 written with 꺼), the edges that go with them (1056 and 1201), and the total of counts and the set of closings are those of the right closings
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
	if bare.graphs != 521 || bare.edges != 1056 || words.graphs != 572 || words.edges != 1201 {
		t.Fatal(bare.graphs, bare.edges, words.graphs, words.edges)
	}
	if bare.closingSet["께"] || bare.closingSet["ㄹ께"] || bare.closingSet["읍니다"] || bare.closingSet["ㄹ게"] != true {
		t.Fatal("the set of closings is those of the tables")
	}
}
