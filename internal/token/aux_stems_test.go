//ff:func feature=token type=reader control=sequence
//ff:what TestAuxStems — the closed set without 지; with a word table 가, 내 and 드리 too
package token

import (
	"slices"
	"testing"
)

func TestAuxStems(t *testing.T) {
	none, with := scanOf("x").ix.auxStems(), wordsScan("x", layerCost).ix.auxStems()
	if !slices.Contains(none, "보") || slices.Contains(none, "지") || slices.Contains(none, "가") {
		t.Fatal(none)
	}
	if !slices.Contains(with, "보") || !slices.Contains(with, "가") || len(with) != len(none)+len(auxMore) || len(none) != 6 {
		t.Fatal(with)
	}
}
