//ff:func feature=token type=reader control=sequence
//ff:what TestFitsParticle — article 41: a boundary after a word that has a particle or the copula after its stem fits; one after a word with a closing or a bare stem does not
package token

import "testing"

func TestFitsParticle(t *testing.T) {
	if !fitsParticle(gap{left: spacingWord("방/열쇠:체언", "에/부사격:에")}) || !fitsParticle(gap{left: spacingWord("총무/열쇠:체언", "이/서술격:이")}) {
		t.Fatal("after a particle")
	}
	if fitsParticle(gap{left: spacingWord("읽/열쇠:동사", "어/맺음:어")}) || fitsParticle(gap{left: spacingWord("방/열쇠:체언")}) {
		t.Fatal("after a closing, a bare stem")
	}
}
