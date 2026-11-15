//ff:func feature=token type=rule control=iteration dimension=2
//ff:what TestMarkerGraphGloss — every marker graph (a particle, a suffix, a slot, a closing, a closing on a lone final consonant, a closing read only with a word table, a closing of clause 40 with 하 dropped) stands for a piece that fits and is turned down for the same piece in a unit glossed with Han characters, with or without a word table
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerGraphGloss(t *testing.T) {
	for _, c := range []struct {
		e    entry
		p    rule.Piece
		word string
	}{
		{entry{name: "목적격:을", form: "을", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkCase}, rule.Piece{Form: "을", HasStem: true, Prev: '곤', Slot: rule.LinkCase, Next: -1}, "대곤을"},
		{entry{name: "관형격:의", form: "의", class: classParticle, slot: rule.LinkAdn}, rule.Piece{Form: "의", HasStem: true, Prev: '치', Slot: rule.LinkAdn, Next: -1}, "관치의"},
		{entry{name: "접미:들", form: "들", class: classParticle, slot: rule.LinkSuffix}, rule.Piece{Form: "들", HasStem: true, Prev: '생', Slot: rule.LinkSuffix, Next: -1}, "학생들"},
		{entry{name: "선어말:었", form: "었", class: classSlot, slot: rule.SlotPast}, rule.Piece{Form: "었", HasStem: true, Prev: '먹', Slot: rule.SlotPast, Next: rule.SlotClosing}, "먹었"},
		{entry{name: "맺음:곤(고는)", form: "곤", class: classClosing}, rule.Piece{Form: "곤", HasStem: true, Prev: '대', Next: -1}, "대곤"},
		{entry{name: "이름:기", form: "기", class: classClosing}, rule.Piece{Form: "기", HasStem: true, Prev: '오', Next: -1}, "오기"},
		{entry{name: "꾸밈:ㄴ", class: classClosing, lead: rule.JongN}, rule.Piece{HasStem: true, Prev: '하', Base: "하", Next: -1}, "한"},
		{entry{name: "맺음:재(자고 해)", form: "재", class: classClosing, wide: true}, rule.Piece{Form: "재", HasStem: true, Prev: '가', Next: -1, Right: rule.RightDone}, "가재"},
		{entry{name: "맺음:지", form: "치", class: classClosing, wide: true}, rule.Piece{Form: "치", HasStem: true, Prev: '관', Next: -1, Right: rule.RightDone}, "관치"},
	} {
		for _, lexed := range []bool{false, true} {
			e := c.e
			markerGraph(&e, nil, nil, lexed)
			if v, _ := verdictOf(e.g, c.p, c.word, false); v <= 0 {
				t.Errorf("%s in %s (word table %v): %v, want it to stand", e.name, c.word, lexed, v)
			}
			glossed := c.p
			glossed.Glossed = true
			v, why := verdictOf(e.g, glossed, c.word, true)
			if v != -1 || len(why) != 1 || why[0] != whyGlossed {
				t.Errorf("%s in %s glossed (word table %v): %v %v", e.name, c.word, lexed, v, why)
			}
		}
	}
}
