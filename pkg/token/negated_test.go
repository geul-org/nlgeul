//ff:func feature=token type=reader control=sequence
//ff:what TestNegated — 했잖아 reads 지+(않) before the closing, alone and after the slot fused into the stem; a tail fused into the letter, a tail of the copula only, another letter, 잖 with nothing before it and a tail at the start give nothing
package token

import "testing"

func TestNegated(t *testing.T) {
	a := []piece{{text: "아"}}
	sc := scanOf("했잖아")
	got := sc.negated([]tail{{start: 2, pieces: a, noYo: true}})
	if len(got) != 2 || got[0].start != 1 || len(got[0].pieces) != 3 || got[0].pieces[0].names != jiName || got[0].pieces[1].text != "(않)" || got[0].close != sc.ix.ji || !got[0].noYo {
		t.Fatal(got)
	}
	if got[1].base != '하' || len(got[1].pieces) != 4 || got[1].close != sc.ix.ji || !got[1].noYo {
		t.Fatal(got[1])
	}
	none := sc.negated([]tail{{start: 2, pieces: a, base: '자'}, {start: 2, pieces: a, only: "아니"}, {start: 1, pieces: a}, {start: 0, pieces: a}})
	if len(none) != 0 {
		t.Fatal(none)
	}
	if got = scanOf("잖아").negated([]tail{{start: 1, pieces: a}}); len(got) != 0 {
		t.Fatal("no stem", got)
	}
}
