//ff:func feature=token type=rule control=sequence
//ff:what TestAddPair — the graph of one candidate of a sign of a pair is put under the sign with the candidate as its form and the sign after a colon in its name, and the graphs and edges are counted
package token

import "testing"

func TestAddPair(t *testing.T) {
	ix := &index{pairs: map[rune][]*entry{}}
	ix.addPair('(', pairRows['('], openName)
	ix.addPair('"', pairRows['"'], shutName)
	if len(ix.pairs['(']) != 1 || ix.pairs['('][0].name != "열기:(" || ix.pairs['('][0].form != openName || ix.pairs['('][0].g == nil {
		t.Fatal(ix.pairs['('])
	}
	if len(ix.pairs['"']) != 1 || ix.pairs['"'][0].name != "닫기:\"" || ix.pairs['"'][0].form != shutName {
		t.Fatal(ix.pairs['"'])
	}
	if ix.graphs != 2 || ix.edges != 1 {
		t.Fatal(ix.graphs, ix.edges)
	}
}
