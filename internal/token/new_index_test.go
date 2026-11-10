//ff:func feature=token type=rule control=sequence
//ff:what TestNewIndex — the index holds the words, particles, closings and slots, and a new word's verdict is the low positive 1/3; with a word table its keys and layers are set, and a table that cannot be read is an error
package token

import (
	"math"
	"testing"
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
}
