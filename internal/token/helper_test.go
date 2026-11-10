//ff:func feature=token type=reader control=sequence
//ff:what TestHelper — the syllable a tail's stem ends in: the one before the tail, or the one a piece was fused into; none at the start, for a tail with no pieces, or for a tail that begins with the copula
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestHelper(t *testing.T) {
	sc := scanOf("지는")
	one := []piece{{text: "는", tail: rule.TailPred}}
	if sc.helper(tail{start: 1, pieces: one}) != '지' || sc.helper(tail{start: 1, pieces: one, base: '하'}) != '하' {
		t.Fatal("the syllable before, or the restored one")
	}
	if sc.helper(tail{start: 0, pieces: one}) != 0 || sc.helper(tail{start: 1}) != 0 || sc.helper(tail{start: 1, pieces: []piece{{text: "이", tail: rule.TailCopula}}}) != 0 {
		t.Fatal("no stem to speak of")
	}
}
