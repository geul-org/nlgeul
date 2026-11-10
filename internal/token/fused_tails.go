//ff:func feature=token type=reader control=iteration dimension=1
//ff:what fusedTails — the tails a piece fused into the stem's last syllable begins: one for each stem the syllable may come from, with the piece written in brackets as it was before fusing; an 어/아 that leaves no trace in the letters, and a lone final consonant under a syllable other than the few it is always read under, ask with a word table for a predicate's stem of the table before them (asksStem)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) fusedTails(pc piece, e *entry, start int) []tail {
	var out []tail
	for _, o := range sc.restoresAt(start, e.lead) {
		p := pc
		p.text = "(" + o.piece + ")" + e.form
		p.pref = sc.ix.fusedPref(e, o.piece, pc.pref)
		p.unseen = e.lead == leadVowel && e.form == "" && (!rule.FusedVowel(sc.s[start-1], false) || o.stem == sc.s[start-1])
		t := tail{start: start, pieces: []piece{p}, base: o.stem, noYo: e.noYo, head: e.head}
		if sc.asksStem(e, start, o.stem) {
			t.gate = sc.ix.pred
		}
		out = append(out, t)
		if o.stem == '시' {
			out = append(out, sc.honors(p, start, e.form == "")...)
		}
	}
	return out
}
