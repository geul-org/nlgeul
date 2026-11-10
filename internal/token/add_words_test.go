//ff:func feature=token type=rule control=sequence
//ff:what TestAddWords — every closed-set word is a key with a graph; a word of two sets (이: 관형사, 대로 only one) keeps both names
package token

import "testing"

func TestAddWords(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addWords()
	if ix.words["그리고"].name != "접속 부사:그리고" || ix.words["것"].g == nil || ix.words["그리고"].class != classKey {
		t.Fatal(ix.words["그리고"])
	}
	if ix.closed["것"] != kindNominal || ix.closed["두"] != kindVerbal|kindDeterminer {
		t.Fatal(ix.closed)
	}
	if both := ix.words["두"].name; both != "관형사:두|보조 용언:두" {
		t.Fatal(both)
	}
}
