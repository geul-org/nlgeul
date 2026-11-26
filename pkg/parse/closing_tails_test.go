//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestClosingTails — a whole closing alone and after slots, a fused one once for each stem it may come from; each carries its closing, and one that attaches only to the copula names 아니, and one right after a digit names a stem nothing is (digitOnly); 죠 is written out as 지+(요); a closing with a gate hands it to its tails; a shortened quotation that only verbs take and the copula takes (란, 래, 랄까) is read once more as a tail of 아니 alone, and one the copula does not take, one of adjectives, a fused one and a closing of the tables are not
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

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
	got = scanOf("3면").closingTails(piece{text: "면"}, &entry{name: "맺음:면", form: "면"}, 1)
	if len(got) == 0 || got[0].only != digitOnly || got[0].start != 1 {
		t.Fatal("a closing right after a digit names a stem nothing is", got)
	}
	got = scanOf("3간").closingTails(piece{}, &entry{name: "꾸밈:ㄴ", lead: 4}, 1)
	if len(got) == 0 || got[0].only == digitOnly {
		t.Fatal("a fused closing is no closing written right after the digit", got)
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
	ran := &entry{name: "꾸밈:란(라는)", form: "란", wide: true, tail: rule.TailVerb, gate: gate}
	got = scanOf("아니란").closingTails(piece{text: "란", tail: rule.TailVerb}, ran, 2)
	last := got[len(got)-1]
	if len(got) != 2 || got[0].only != "" || got[0].pieces[0].tail != rule.TailVerb || last.only != notOnly || last.pieces[0].tail != rule.TailPred || last.gate != gate || last.close != ran || last.start != 2 {
		t.Fatal("a shortened quotation only verbs take is read once more for 아니", got)
	}
	for _, e := range []*entry{
		{name: "꾸밈:잔(자는)", form: "잔", wide: true, tail: rule.TailVerb, bars: true},
		{name: "꾸밈:단(다는)", form: "단", wide: true, tail: rule.TailAdj},
		{name: "꾸밈:ㄴ단(ㄴ다는)", form: "단", lead: rule.JongN, wide: true, tail: rule.TailVerb},
		{name: "맺음:느라", form: "란", tail: rule.TailVerb},
	} {
		for _, x := range scanOf("아니란").closingTails(piece{text: "란"}, e, 2) {
			if x.only != "" {
				t.Fatal("no tail for 아니", e.name, x)
			}
		}
	}
}
