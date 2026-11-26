//ff:func feature=parse type=reader control=selection
//ff:what politeClosing — the closing of a tail as it stands right after an old polite slot that ends in last (오 or 옵), and whether it stands there: a tail of one closing alone; a slot in 오 before a closing that begins with ㄴ, ㄹ, ㅁ or a vowel (also one fused under 오: 온, 올, 옴), a slot in 옵 before a closing that begins with a consonant, and neither where the slot's last syllable and the closing are together a closing of the tables (오니, 오니까, 오리다: the tables read those); ㅂ니다 fused under 오 is 옵 and the closing 니다
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) politeClosing(t tail, last rune) (piece, bool) {
	switch {
	case t.only != "" || t.shown != "" || t.as != "" || t.irr != 0 || t.start < 1 || t.close == nil || t.close.wrote != "" || len(t.pieces) != 1 || t.pieces[0].names != t.close.name:
		return piece{}, false
	case t.base == 0:
		return t.pieces[0], sc.s[t.start-1] == last && t.start < len(sc.s) && politeBefore(last, sc.s[t.start]) && !sc.ix.closingSet[string(last)+t.close.form]
	case t.base != politeSoft:
		return piece{}, false
	case last == politeHard:
		c := t.pieces[0]
		c.text, c.names = politeFinal, closingLabel(0)+":"+politeFinal
		return c, t.close.lead == rule.JongB && t.close.form == politeFinal
	}
	return t.pieces[0], t.close.lead == rule.JongN || t.close.lead == rule.JongL || t.close.lead == rule.JongM
}
