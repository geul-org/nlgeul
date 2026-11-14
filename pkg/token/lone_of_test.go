//ff:func feature=token type=reader control=sequence
//ff:what TestLoneOf — a whole unit read as a tail alone is a word with no stem part, marked as attaching to the unit before and printed with a leading "+"; its verdict is no higher than a new word's nor than its pieces', its preference their sum less the structure of a word with an unseen stem (and each unseen piece)
package token

import "testing"

func TestLoneOf(t *testing.T) {
	sc := scanOf("인데")
	cop := tail{start: 0, pieces: []piece{
		{text: "이", names: "서술격:이", verdict: 1, pref: -1},
		{text: "(ㄴ)데", names: "맺음:ㄴ데", verdict: 1, pref: -2},
	}}
	e := sc.loneOf(cop, 2)
	if e.end != 2 || e.word.End != 2 || !e.word.Attach || len(e.word.Parts) != 2 {
		t.Fatal(e)
	}
	if p := e.word.Parts[0]; p.Stem || p.Text != "이" || p.Hidden != "" || p.Name != (Name{"서술격", "이"}) {
		t.Fatal(p)
	}
	if p := e.word.Parts[1]; p.Stem || p.Text != "데" || p.Hidden != "ㄴ" || p.Name != (Name{"맺음", "ㄴ데"}) {
		t.Fatal(p)
	}
	if e.verdict != sc.ix.unknown || e.pref != -(costPiece+costUnseen)-1-2 || e.sig != "+이[서술격:이]+(ㄴ)데[맺음:ㄴ데]" || e.sig != e.word.String() {
		t.Fatal(e.verdict, e.pref, e.sig)
	}
	low := sc.ix.unknown / 2
	drop := tail{start: 0, pieces: []piece{
		{text: "(이)", names: "서술격:이(빠짐)", verdict: low, pref: -1, unseen: true},
		{text: "다", names: "맺음:다", verdict: 1, pref: -1},
	}}
	e = scanOf("다").loneOf(drop, 1)
	if e.verdict != low || e.pref != -(costPiece+costUnseen+costUnseen)-2 || !e.word.Parts[0].Unseen || e.word.Parts[0].Hidden != "이" || e.word.Parts[0].Text != "" || e.sig != "+(이)[서술격:이(빠짐)]+다[맺음:다]" {
		t.Fatal(e.verdict, e.pref, e.sig)
	}
}
