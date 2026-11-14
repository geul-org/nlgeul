//ff:func feature=token type=helper control=sequence
//ff:what TestHasParticleOnly — a particle among the parts after the stem; the copula is not one
package token

import "testing"

func TestHasParticleOnly(t *testing.T) {
	if !hasParticleOnly(spacingWord("듯/열쇠:체언", "도/보조사:도")) {
		t.Fatal("a particle")
	}
	if hasParticleOnly(spacingWord("총무/열쇠:체언", "이/서술격:이")) || hasParticleOnly(spacingWord("총무/열쇠:체언")) {
		t.Fatal("none")
	}
}
