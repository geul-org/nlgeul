//ff:func feature=token type=rule control=sequence
//ff:what restored — may the stem from b have the last syllable tail t reads it with, a piece taken out of the syllable before the tail: the graph of a restored syllable is told the syllable, the letter of the stem before it (none at the stem's head), and for a ㅡ put back under ㅏ or ㅓ (따라 = 따르+아) the 어/아 that was fused in and whether the word table gives the stem as 르 or 러 irregular; a tail fused into nothing is not asked
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) restored(b int, t tail) bool {
	if t.base == 0 {
		return true
	}
	p := rule.Piece{Base: string(t.base), HasStem: t.start-2 >= b}
	if p.HasStem {
		p.Prev = sc.s[t.start-2]
	}
	shown := rule.JungOf(sc.s[t.start-1])
	if rule.JungOf(t.base) == jungEu && (shown == jungA || shown == jungEo) {
		stem := string(sc.s[b:t.start-1]) + string(t.base)
		p.Form, p.Keeps = map[int]string{jungA: "아", jungEo: "어"}[shown], sc.ix.lex[stem].irr&(irrReu|irrReo) != 0
	}
	return verdict(sc.ix.restore, p) > 0
}
