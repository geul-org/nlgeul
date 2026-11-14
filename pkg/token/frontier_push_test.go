//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestFrontierPush — a walk goes in behind the walks of f not less than its own: the greatest f first, equal f in the order put in, also minus infinity; the most walks held at one time is kept (peak)
package token

import (
	"math"
	"testing"
)

func TestFrontierPush(t *testing.T) {
	var f frontier
	inf := math.Inf(-1)
	for i, v := range []float64{-5, inf, -2, -5, inf, -1} {
		f.push(&walk{node: i, f: v})
	}
	got := ""
	for _, w := range f.items {
		got += string(rune('0' + w.node))
	}
	if got != "520314" || f.peak != 6 {
		t.Fatal(got, f.peak)
	}
	f.pop()
	f.push(&walk{node: 9, f: -3})
	if f.peak != 6 || len(f.items) != 6 {
		t.Fatal("the most at one time", f.peak)
	}
}
