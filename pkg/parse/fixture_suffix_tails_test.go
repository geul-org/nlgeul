//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what checkSuffixTails — report as test errors the tails that do not hold one of the suffixes and particles of copulaBefore right before the copula stem and the closing, or that lost the closing, noYo or the gate
package parse

import (
	"slices"
	"testing"
)

func checkSuffixTails(t *testing.T, text string, got []tail, e, g *entry) {
	t.Helper()
	for _, r := range got {
		last := r.pieces[len(r.pieces)-3].names
		if !slices.Contains(copulaBefore, last) || r.close != e || !r.noYo || r.gate != g || r.pieces[len(r.pieces)-1].text != "다" {
			t.Errorf("%s: %+v", text, r)
		}
	}
}
