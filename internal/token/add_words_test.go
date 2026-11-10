//ff:func feature=token type=rule control=sequence
//ff:what TestAddWords — every closed-set word is a key with a graph; a word of two sets (이: 관형사, 대로 only one) keeps both names
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestAddWords(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addWords()
	if ix.words["그리고"].name != "접속 부사:그리고" || ix.words["것"].g == nil || ix.words["그리고"].class != classKey {
		t.Fatal(ix.words["그리고"])
	}
	if ix.closed["것"] != kindNominal || ix.closed["두"] != kindVerb|kindAdj|kindDeterminer {
		t.Fatal(ix.closed)
	}
	if both := ix.words["두"].name; both != "관형사:두|보조 용언:두" {
		t.Fatal(both)
	}
	if ix.graphs != len(ix.words) || ix.edges != len(ix.words) || ix.takes["두"]&rule.TailPred == 0 || ix.takes["두"]&rule.TailBare == 0 || ix.takes["주"]&rule.TailBare != 0 {
		t.Fatal(ix.graphs, ix.edges, ix.takes)
	}
}
