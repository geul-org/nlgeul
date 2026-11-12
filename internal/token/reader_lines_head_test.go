//ff:func feature=token type=reader control=sequence
//ff:what TestReaderLinesHead — a contest of more candidates than head prints its first head candidates and the count left out; the trace lines are not cut; Lines prints all; and Bundles gives the same candidates whatever head is
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestReaderLinesHead(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	s := "20m²"
	all := r.Lines(s, false)
	n := len(r.Bundles(s)[0].Branches)
	if n < 4 || len(all) != n+1 || slices.Contains(all, "  (1개 생략됨)") {
		t.Fatal(n, all)
	}
	cut := r.LinesHead(s, false, 3)
	if len(cut) != 5 || cut[4] != "  ("+string(rune('0'+n-3))+"개 생략됨)" || !slices.Equal(cut[:4], all[:4]) {
		t.Fatal(cut)
	}
	if exact := r.LinesHead(s, false, n); !slices.Equal(exact, all) {
		t.Fatal(exact)
	}
	if zero := r.LinesHead(s, false, 0); !slices.Equal(zero, all) {
		t.Fatal(zero)
	}
	if got := len(r.Bundles(s)[0].Branches); got != n {
		t.Fatal(got)
	}
	traced := r.LinesHead(s, true, 3)
	if len(traced) <= len(cut) || strings.Count(strings.Join(traced, "\n"), "생략됨") != 1 {
		t.Fatal(traced)
	}
}
