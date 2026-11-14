//ff:func feature=token type=reader control=sequence
//ff:what TestClosingTails — a whole closing alone and after slots, a fused one once for each stem it may come from; each carries its closing, and one that attaches only to the copula names 아니; 죠 is written out as 지+(요); a closing with a gate hands it to its tails
package token

import "testing"

func TestClosingTails(t *testing.T) {
	sc := scanOf("먹었다")
	e := &entry{name: "맺음:다", form: "다", noYo: true}
	got := sc.closingTails(piece{text: "다"}, e, 2)
	if len(got) < 2 || got[0].start != 2 || got[1].start != 1 || got[1].close != e || !got[1].noYo || got[0].only != "" {
		t.Fatal(got)
	}
	f := &entry{name: "꾸밈:ㄴ", lead: 4}
	got = scanOf("간").closingTails(piece{}, f, 1)
	if len(got) == 0 || got[0].base != '가' || got[0].close != f {
		t.Fatal(got)
	}
	got = sc.closingTails(piece{text: "에요"}, &entry{name: "맺음:에요", needs: true}, 3)
	if len(got) != 1 || got[0].only != notOnly {
		t.Fatal(got)
	}
	got = scanOf("하죠").closingTails(piece{text: "죠"}, &entry{name: "맺음:죠", form: "죠"}, 1)
	if len(got) != 1 || len(got[0].pieces) != 2 || got[0].pieces[1].text != "(요)" {
		t.Fatal(got)
	}
	gate := &entry{name: "g"}
	got = sc.closingTails(piece{text: "다"}, &entry{name: "맺음:다", form: "다", gate: gate}, 2)
	if got[0].gate != gate || got[1].gate != gate {
		t.Fatal(got)
	}
}
