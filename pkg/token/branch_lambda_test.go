//ff:func feature=token type=helper control=sequence
//ff:what TestBranchLambda — a candidate's lambda is its preference plus Lambda times log10 of its verdict (a verdict of 1 adds nothing, 0.1 takes Lambda away), minus infinity for a verdict of 0 or less
package token

import (
	"math"
	"testing"
)

func TestBranchLambda(t *testing.T) {
	if got := (Branch{Verdict: 1, Pref: -2}).Lambda(); got != -2 {
		t.Fatal(got)
	}
	if got := (Branch{Verdict: 0.1, Pref: -2}).Lambda(); math.Abs(got-(-2-Lambda)) > 1e-9 {
		t.Fatal(got)
	}
	if got := (Branch{Verdict: 0, Pref: 3}).Lambda(); !math.IsInf(got, -1) {
		t.Fatal(got)
	}
}
