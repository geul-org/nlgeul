//ff:func feature=token type=helper control=sequence
//ff:what TestWayTag — the rank in a circle and the lambda with two decimals
package token

import "testing"

func TestWayTag(t *testing.T) {
	if got := wayTag(Segment{Rank: 2, Lambda: -12.366}); got != "② λ 합 -12.37" {
		t.Fatal(got)
	}
	if got := wayTag(Segment{Rank: 1, Lambda: 0}); got != "① λ 합 0.00" {
		t.Fatal(got)
	}
}
