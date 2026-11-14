//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestRestOf — what a way can reach from a node on: its best lambda plus the greatest of the nodes after it (its own alone at the end); a node that only leads to minus infinity reaches minus infinity
package token

import (
	"math"
	"testing"
)

func TestRestOf(t *testing.T) {
	best := []float64{-1, -2, -5, -3, -4}
	next := [][]int{{1, 2}, {3}, {4}, nil, nil}
	got := restOf(best, next)
	want := []float64{-1 + -2 + -3, -5, -9, -3, -4}
	for i := range want {
		if got[i] != want[i] {
			t.Fatal(got)
		}
	}
	got = restOf([]float64{-1, math.Inf(-1)}, [][]int{{1}, nil})
	if !math.IsInf(got[0], -1) || !math.IsInf(got[1], -1) {
		t.Fatal(got)
	}
}
