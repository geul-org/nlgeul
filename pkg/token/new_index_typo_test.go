//ff:func feature=token type=rule control=sequence
//ff:what TestNewIndexTypo — the index has a graph for each closing written wrong (17 more than the closings alone) and one for the family names (addSurname): 710 without a word table, 780 with one, which also reads two forms of 거 written with 꺼, the six forms of clause 40 with 하 dropped and the suffix 률 with the graph of the word before it; the edges that go with them (1999 and 2237: each marker graph but a suffix's has the exception of a unit glossed with Han characters, 5 suffixes without a word table and 6 with one (률), the graph of the naming closing 이름:ㅁ has the exception of its row in the case exception table, 고름, and with a word table the graph of the 하오체 closing 맺음:소 has that of its row, 담소). The dependent noun added to the closed set after 만큼 (간) is one graph and one edge more (a key of its own, with the exception that its set does not take what follows it), and the total of counts and the set of closings are those of the right closings
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
	if bare.graphs != 710 || bare.edges != 1999 || words.graphs != 780 || words.edges != 2237 {
		t.Fatal(bare.graphs, bare.edges, words.graphs, words.edges)
	}
	if bare.closingSet["께"] || bare.closingSet["ㄹ께"] || bare.closingSet["읍니다"] || bare.closingSet["ㄹ게"] != true {
		t.Fatal("the set of closings is those of the tables")
	}
}
