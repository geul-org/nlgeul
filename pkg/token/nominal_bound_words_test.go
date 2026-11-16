//ff:func feature=token type=rule control=sequence
//ff:what TestNominalBound — the words nominalBoundRule asks are the dependent nouns of the table that stand after a nominal and bear the word table's name (때문, 중, 내, 측, 분), not those of the closed set that stand there too (간, 외, 때, 시) nor those that stand after a modifying closing only (리, 자)
package token

import (
	"slices"
	"testing"
)

func TestNominalBound(t *testing.T) {
	in := func(w string) bool { return slices.Contains(nominalBound, w) }
	if !in("때문") || !in("중") || !in("내") || !in("측") || !in("분") {
		t.Errorf("missing a word of the table: %v", nominalBound)
	}
	if in("간") || in("외") || in("때") || in("시") || in("리") || in("자") {
		t.Errorf("a word of the closed set or after a modifying closing only: %v", nominalBound)
	}
}
