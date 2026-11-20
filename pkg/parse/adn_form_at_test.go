//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAdnFormAt — the letters read as the modifying form of a predicate: 할 (하+ㄹ), 먹는, 작은 (an adjective); not a noun (학교), not a predicate with a closing that is no modifying one (먹고), not a modifying closing after a stem the tables do not know (곰는), not ㄴ fused into a syllable right after a Latin letter (x간: the syllable is not restored there)
package parse

import "testing"

func TestAdnFormAt(t *testing.T) {
	for _, c := range []struct {
		text string
		want bool
	}{
		{"할", true},
		{"먹는", true},
		{"작은", true},
		{"학교", false},
		{"먹고", false},
		{"곰는", false},
		{"x간", false},
	} {
		sc := wordsScan(c.text, layerCost)
		if got := sc.adnFormAt(0, len(sc.s)); got != c.want {
			t.Errorf("%s: %v, want %v", c.text, got, c.want)
		}
	}
}
