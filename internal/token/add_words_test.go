//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddWords — every closed-set word is a key with a graph; a word of two sets (이: 관형사, 대로 only one) keeps both names; this layer's pronouns 네, 내 and 제 are nominals in the closed kinds but no keys, so that with a word table each stands before a particle (네+가) and without one none is read
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddWords(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addWords()
	if ix.words["그리고"].name != "접속 부사:그리고" || ix.words["것"].g == nil || ix.words["그리고"].class != classKey {
		t.Fatal(ix.words["그리고"])
	}
	if ix.closed["것"] != kindNominal || ix.closed["두"] != kindVerb|kindDeterminer {
		t.Fatal(ix.closed)
	}
	if both := ix.words["두"].name; both != "관형사:두|보조 용언:두" {
		t.Fatal(both)
	}
	if ix.graphs != len(ix.words) || ix.edges != len(ix.words) || ix.takes["두"]&rule.TailPred == 0 || ix.takes["두"]&rule.TailBare == 0 || ix.takes["주"]&rule.TailBare != 0 {
		t.Fatal(ix.graphs, ix.edges, ix.takes)
	}
	for _, w := range []string{"네", "내", "제"} {
		if _, key := ix.words[w]; key || ix.closed[w]&kindNominal == 0 {
			t.Errorf("%s: %v %d", w, key, ix.closed[w])
		}
		if _, v := wordsScan(w, layerCost).ix.known(w, rule.TailCase, false); v <= 0 {
			t.Errorf("%s before a particle with a word table", w)
		}
	}
	r, _ := New(Options{Words: "testdata/words.tsv"})
	if got := strings.Join(r.Lines("네가 내가", false), "\n"); !strings.Contains(got, "네[열쇠:체언]+가[주격:가|보격:가]") || !strings.Contains(got, "내[열쇠:체언]+가[주격:가|보격:가]") {
		t.Fatal("네가 내가\n" + got)
	}
	bare, _ := New(Options{})
	if got := strings.Join(bare.Lines("네가", false), "\n"); strings.Contains(got, "네[") {
		t.Fatal("네 without a word table\n" + got)
	}
}
