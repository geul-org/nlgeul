//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestParticlePart — a case particle, the vocative and an auxiliary particle are particles; the copula and a closing are not
package token

import "testing"

func TestParticlePart(t *testing.T) {
	for _, kind := range []string{"주격", "보격", "목적격", "관형격", "부사격", "호격", "보조사"} {
		if !particlePart(Part{Name: Name{Kind: kind}}) {
			t.Fatal(kind)
		}
	}
	for _, kind := range []string{"서술격", "맺음", "열쇠", ""} {
		if particlePart(Part{Name: Name{Kind: kind}}) {
			t.Fatal(kind)
		}
	}
}
