//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAfterSuffixes — a copula tail is read once more after the suffix 들, 씩 or 끼리 that ends where it starts and keeps its closing and its gate; after 쯤 or the suffix 이 of a name nothing is read, nor for a tail fused into a stem or one with no suffix before it
package token

import "testing"

func TestAfterSuffixes(t *testing.T) {
	e := &entry{name: "맺음:다"}
	g := &entry{name: "앞말:체언"}
	tl := tail{start: 3, pieces: []piece{{text: "이", names: copulaName}, {text: "다", names: "맺음:다"}}, close: e, noYo: true, gate: g}
	for _, text := range []string{"곰들이다", "곰씩이다", "곰끼리이다"} {
		sc := scanOf(text)
		tl.start = len([]rune(text)) - 2
		got := sc.afterSuffixes(tl)
		if len(got) == 0 {
			t.Fatalf("%s: nothing read", text)
		}
		checkSuffixTails(t, text, got, e, g)
	}
	if got := scanOf("곰쯤이다").afterSuffixes(tail{start: 2, pieces: tl.pieces}); len(got) != 0 {
		t.Fatal("after 쯤", got)
	}
	if got := scanOf("영숙이이다").afterSuffixes(tail{start: 3, pieces: tl.pieces}); len(got) != 0 {
		t.Fatal("after the suffix 이", got)
	}
	if got := scanOf("곰들이다").afterSuffixes(tail{start: 2, pieces: tl.pieces, base: '가'}); len(got) != 0 {
		t.Fatal("a tail fused into a stem", got)
	}
	if got := scanOf("이다").afterSuffixes(tail{start: 0, pieces: tl.pieces}); len(got) != 0 {
		t.Fatal("no suffix before the tail", got)
	}
}
