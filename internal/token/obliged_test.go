//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestObliged — 겠 right after 어야 (가야겠다) is read once more after each candidate of that closing that ends before it, with (하) written out between: the closing keeps its start and its stem's restored syllable, (하) takes 겠's verdict and preference, 겠 keeps its verdict with none, and the candidate takes 요 and the closing of the slot's tail; so too after 어야, 아야 and 여야 written whole and after a slot before 어야 (갔어야겠지); a 야 that attaches only to the copula, a closing that is not 어야, and a tail fused into the letter, with its own stem, shown word, name or change, at the start or right after it, of 겠 alone or of another slot give nothing
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestObliged(t *testing.T) {
	gyet := piece{"겠", guessName, 0.5, -3, rule.TailPred, false, ""}
	da := piece{"다", "맺음:다", 1, -1, rule.TailPred, false, ""}
	end := &entry{name: "맺음:다"}
	got := scanOf("가야겠다").obliged([]tail{{start: 2, pieces: []piece{gyet, da}, noYo: true, close: end}})
	if len(got) != 2 || got[0].base != '가' || got[1].base != '그' {
		t.Fatal("one candidate for each stem the fused 어야 may come from, and none after the copula's 야", got)
	}
	for _, l := range got {
		ps := l.pieces
		if l.start != 1 || !l.noYo || l.close != end || len(ps) != 4 || ps[0].text != "(아)야" || ps[0].names != "맺음:어야(합침)" || ps[3] != da {
			t.Fatal(l)
		}
		if ps[1] != (piece{"(하)", "보조:하", 0.5, -3, rule.TailPred, false, ""}) || ps[2] != (piece{"겠", guessName, 0.5, 0, rule.TailPred, false, ""}) {
			t.Fatal("(하) between the closing and 겠, with 겠's preference", ps)
		}
	}
	whole := []struct {
		text  string
		start int
		name  string
	}{
		{"먹어야겠네", 3, "맺음:어야"},
		{"잡아야겠다", 3, "맺음:아야"},
		{"하여야겠다", 3, "맺음:여야"},
	}
	for _, c := range whole {
		got = scanOf(c.text).obliged([]tail{{start: c.start, pieces: []piece{gyet, da}, close: end}})
		if len(got) == 0 || got[0].start != 1 || got[0].base != 0 || got[0].pieces[0].names != c.name || got[0].pieces[1].names != "보조:하" || got[0].pieces[2].names != guessName {
			t.Errorf("%s: %v", c.text, got)
		}
	}
	slot := scanOf("갔어야겠지").obliged([]tail{{start: 3, pieces: []piece{gyet, {"지", "맺음:지", 1, -1, rule.TailPred, false, ""}}, close: end}})
	past := func(l tail) bool {
		return l.start == 1 && l.base == '가' && len(l.pieces) == 5 && l.pieces[0].names == "선어말:ㅆ" && l.pieces[1].names == "맺음:어야" && l.pieces[2].text == "(하)"
	}
	if !slices.ContainsFunc(slot, past) {
		t.Fatal("a slot before 어야", slot)
	}
	none := []struct {
		text string
		t    tail
	}{
		{"가야겠다", tail{start: 2, pieces: []piece{gyet, da}, base: '게'}},
		{"가야겠다", tail{start: 2, pieces: []piece{gyet, da}, only: notOnly}},
		{"가야겠다", tail{start: 2, pieces: []piece{gyet, da}, shown: "가야"}},
		{"가야겠다", tail{start: 2, pieces: []piece{gyet, da}, as: numName}},
		{"가야겠다", tail{start: 2, pieces: []piece{gyet, da}, irr: irrL}},
		{"가야겠다", tail{start: 2, pieces: []piece{gyet}}},
		{"가야겠다", tail{start: 2, pieces: []piece{{"었", "선어말:었", 1, -1, rule.TailPred, false, ""}, da}}},
		{"야겠다", tail{start: 1, pieces: []piece{gyet, da}}},
		{"먹고겠다", tail{start: 2, pieces: []piece{gyet, da}}},
	}
	for i, c := range none {
		if got = scanOf(c.text).obliged([]tail{c.t}); len(got) != 0 {
			t.Errorf("%d %s: %v", i, c.text, got)
		}
	}
}
