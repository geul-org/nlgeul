//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCutMarked — the Hangul after the slash a path is cut at reads as markers on the path when a run of markers that ends at the unit's end starts at the cut (a particle, a chain of particles, the copula with a closing fused into it); a word, and a word with a particle after it, do not; a unit with no cut has none, and prepare notes the answer (onSlash)
package parse

import "testing"

func TestCutMarked(t *testing.T) {
	for in, want := range map[string]bool{"~/a/b/에": true, "~/a/b/에서는": true, "~/a/b/인": true, "~/a/b/회의록": false, "~/a/b/회의록을": false, "~/a/b를": false, "고양이가": false} {
		sc := scanOf(in)
		if got := sc.cutMarked(); got != want || sc.onSlash != want {
			t.Errorf("%s: %v %v (cut %d)", in, got, sc.onSlash, sc.cut)
		}
	}
}
