//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestNumeralBeforeSame — numeralBefore gives the answer of numeralBeforeAll (the same question asked back to the unit's head) at every place of thousands of made-up units: random letters of the numerals, of the ones that stand beside them (수, 열, 한, 두), and of others
package token

import (
	"math/rand"
	"testing"
)

func TestNumeralBeforeSame(t *testing.T) {
	alphabet := []rune("일이삼사오육칠팔구십백천만억조수열한두세네다섯여섯곱덟아홉스무몇러개원가")
	rng := rand.New(rand.NewSource(7))
	yes := 0
	for i := 0; i < 3000; i++ {
		text := make([]rune, 1+rng.Intn(14))
		for k := range text {
			text[k] = alphabet[rng.Intn(len(alphabet))]
		}
		sc := &scan{s: text}
		for p := 0; p <= len(text); p++ {
			want := numeralBeforeAll(sc, p)
			if got := sc.numeralBefore(p); got != want {
				t.Fatal(string(text), p, got, want)
			}
			if want {
				yes++
			}
		}
	}
	if yes < 1000 {
		t.Fatal("too few places that read as numerals", yes)
	}
}
