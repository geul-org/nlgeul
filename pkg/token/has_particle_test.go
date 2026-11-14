//ff:func feature=token type=helper control=sequence
//ff:what TestHasParticle — a particle or the copula among the parts after the stem; none when there is none (the stem itself does not count)
package token

import "testing"

func TestHasParticle(t *testing.T) {
	if !hasParticle(spacingWord("씨/열쇠:체언", "가/주격:가")) || !hasParticle(spacingWord("총무/열쇠:체언", "이/서술격:이", "다/맺음:다")) {
		t.Fatal("a particle, the copula")
	}
	if hasParticle(spacingWord("읽/열쇠:동사", "어/맺음:어")) || hasParticle(spacingWord("가/주격:가")) {
		t.Fatal("none")
	}
}
