//ff:func feature=token type=reader control=sequence
//ff:what TestRank — a reading's place is its preference plus log10 of its verdict
package token

import (
	"math"
	"testing"
)

func TestRank(t *testing.T) {
	if got := rank(&path{verdict: 0.1, pref: -3}); math.Abs(got+4) > 1e-9 {
		t.Fatal(got)
	}
}
