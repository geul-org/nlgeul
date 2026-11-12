//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNewIndex — the index holds the words, particles, closings and slots, and a new word's verdict is the low positive 1/3; with a word table its keys and layers are set, and a table that cannot be read is an error; the new word's graph turns down a stem that ends in ㅆ with markers after it (떨어졌 before a predicate tail or particles) and keeps 1/3 for it with no marker (TailBare or none) and for a stem of another final consonant or none (곰갓, 떨어지)
package token

import (
	"math"
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNewIndex(t *testing.T) {
	withWords, err := newIndex(Options{Words: "testdata/words.tsv"})
	if err != nil || withWords.layers != layerCost || len(withWords.lex) == 0 || withWords.kinds == nil {
		t.Fatal(err)
	}
	if two, _ := newIndex(Options{Words: "testdata/words.tsv", Layers: layerKind}); two.layers != layerKind {
		t.Fatal(two.layers)
	}
	if _, err := newIndex(Options{Words: "testdata/none.tsv"}); err == nil {
		t.Fatal("a missing word table accepted")
	}
	ix, _ := newIndex(Options{})
	if ix.graphs < 200 || ix.total <= 0 || len(ix.words) == 0 || ix.stem == nil || math.Abs(ix.unknown-1.0/3) > 1e-9 {
		t.Fatal(ix.graphs, ix.total, ix.unknown)
	}
	if v, why := verdictOf(ix.stem, rule.Piece{Form: "떨어졌", Tail: rule.TailPred}, ""); v > 0 || !slices.Contains(why, reasons["StemCoda"]) {
		t.Fatal("떨어졌 before a predicate tail", v, why)
	}
	if v, _ := verdictOf(ix.stem, rule.Piece{Form: "떨어졌", Tail: rule.TailCase}, ""); v > 0 {
		t.Fatal("떨어졌 before particles", v)
	}
	for _, p := range []rule.Piece{{Form: "떨어졌", Tail: rule.TailBare}, {Form: "떨어졌"}, {Form: "떨어지", Tail: rule.TailPred}, {Form: "곰갓", Tail: rule.TailPred}} {
		if v, why := verdictOf(ix.stem, p, ""); math.Abs(v-1.0/3) > 1e-9 || len(why) != 0 {
			t.Fatal(p, v, why)
		}
	}
}
