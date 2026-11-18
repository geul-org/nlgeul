//ff:func feature=token type=rule control=sequence
//ff:what restored — may the stem from b have the last syllable tail t reads it with, a piece taken out of the syllable before the tail: the graph of a restored syllable is told the syllable, whether it stands in the letters as it is (the piece left no trace), the letter of the stem before it (none at the stem's head), and for a ㅡ put back under ㅏ or ㅓ (따라 = 따르+아) the 어/아 that was fused in and whether the word table gives the stem as 르 or 러 irregular, unless the tail reads the letters as the changed form of another stem of the table (올랐 = 오르 → 올르+았: the change itself drops the ㅡ, whatever the table says of 올르); a tail fused into nothing is not asked
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// A tail that shows another stem and keeps no letters of a wrong form
// (changedCandidates) reads the letters as that stem's changed form:
// the stem looked up there is the one it shows, and the ㅡ is the one
// the change of 르 or 러 drops before 어/아 (Hangul orthography, clause
// 18, 9 and 8), so the stem of the letters keeping its ㅡ is not what is
// read.
func (sc *scan) restored(b int, t tail) bool {
	if t.base == 0 {
		return true
	}
	p := rule.Piece{Base: string(t.base), HasStem: t.start-2 >= b, Same: sc.s[t.start-1] == t.base}
	if p.HasStem {
		p.Prev = sc.s[t.start-2]
	}
	shown := rule.JungOf(sc.s[t.start-1])
	if rule.JungOf(t.base) == jungEu && (shown == jungA || shown == jungEo) {
		stem := string(sc.s[b:t.start-1]) + string(t.base)
		p.Form, p.Keeps = map[int]string{jungA: "아", jungEo: "어"}[shown], (t.shown == "" || t.wrote != "") && sc.ix.lex[stem].irr&(irrReu|irrReo) != 0
	}
	return sc.ix.judged(sc.ix.restore, p, "") > 0
}
