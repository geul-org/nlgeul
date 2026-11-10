//ff:func feature=token type=rule control=sequence
//ff:what TestNewIndex — the index holds the words, particles, closings and slots, and a new word's verdict is the low positive 1/3
package token

import (
	"math"
	"testing"
)

func TestNewIndex(t *testing.T) {
	ix := newIndex()
	if ix.graphs < 200 || ix.total <= 0 || len(ix.words) == 0 || ix.stem == nil || math.Abs(ix.unknown-1.0/3) > 1e-9 {
		t.Fatal(ix.graphs, ix.total, ix.unknown)
	}
}
