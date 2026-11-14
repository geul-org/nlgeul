//ff:func feature=token type=reader control=sequence
//ff:what TestLoneOfTypo — a unit read as markers alone, one of them written wrong, keeps the letters it is written with in its part and is a candidate with letters written wrong
package token

import "testing"

func TestLoneOfTypo(t *testing.T) {
	sc := scanOf("예요")
	tl := tail{start: 0, pieces: []piece{{text: "에요", names: "맺음:에요", verdict: 1, wrote: "예요"}}}
	e := sc.loneOf(tl, 2)
	if e.word.Parts[0].Written != "예요" || e.sig != "+에요[맺음:에요,오타정정:예요]" {
		t.Fatal(e.sig)
	}
	tl.pieces[0].wrote = ""
	if e = sc.loneOf(tl, 2); e.sig != "+에요[맺음:에요]" {
		t.Fatal(e.sig)
	}
}
