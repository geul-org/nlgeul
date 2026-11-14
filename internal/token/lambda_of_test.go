//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestLambdaOf — a candidate's lambda is its preference plus Lambda times log10 of its verdict (the preference alone at verdict 1), and minus infinity at a verdict of 0 or less
package token

import (
	"math"
	"testing"
)

func TestLambdaOf(t *testing.T) {
	if got := lambdaOf(Branch{Verdict: 1, Pref: -3}); got != -3 {
		t.Fatal(got)
	}
	if got, want := lambdaOf(Branch{Verdict: 0.5, Pref: -3}), -3+15*math.Log10(0.5); math.Abs(got-want) > 1e-12 {
		t.Fatalf("%v, want %v", got, want)
	}
	for _, v := range []float64{0, -1} {
		if got := lambdaOf(Branch{Verdict: v, Pref: -3}); !math.IsInf(got, -1) {
			t.Fatalf("verdict %v: %v", v, got)
		}
	}
	if Lambda != 15 {
		t.Fatal(Lambda)
	}
}
