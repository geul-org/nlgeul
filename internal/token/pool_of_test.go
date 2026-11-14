//ff:func feature=token type=helper control=sequence
//ff:what TestPoolOf — the words at a span are gathered from all the ways, a word printed the same once (the first kept); a word printed the same that came with other values is a clash, counted once however often it comes
package token

import "testing"

func TestPoolOf(t *testing.T) {
	x := streamCut([]Span{{0, 3}, {3, 6}}, []string{"가", "나"}, []float64{-1, -2})
	y := streamCut([]Span{{0, 3}, {3, 6}}, []string{"가", "다"}, []float64{-1, -2})
	z := streamCut([]Span{{0, 3}, {3, 6}}, []string{"가", "나"}, []float64{-9, -2})
	u := streamCut([]Span{{0, 3}, {3, 6}}, []string{"가", "나"}, []float64{-7, -2})
	p := poolOf([]cut{x, y, z, u})
	if got := p.words[Span{0, 3}]; len(got) != 1 || got[0].Pref != -1 {
		t.Fatal(got)
	}
	if got := p.words[Span{3, 6}]; len(got) != 2 || got[0].String() != "나" || got[1].String() != "다" {
		t.Fatal(got)
	}
	if p.clashes != 1 {
		t.Fatal(p.clashes)
	}
	if q := poolOf([]cut{x, y}); q.clashes != 0 || len(q.words) != 2 {
		t.Fatal(q)
	}
}
