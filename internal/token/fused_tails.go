//ff:func feature=token type=reader control=iteration dimension=1
//ff:what fusedTails — the tails a piece fused into the stem's last syllable begins: one for each stem the syllable may come from, with the piece written in brackets as it was before fusing
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) fusedTails(pc piece, e *entry, start int) []tail {
	var out []tail
	for _, o := range restores(sc.s[start-1], e.lead) {
		p := pc
		p.text = "(" + o.piece + ")" + e.form
		p.unseen = e.lead == leadVowel && e.form == "" && !rule.FusedVowel(sc.s[start-1], false)
		out = append(out, tail{start: start, pieces: []piece{p}, base: o.stem, noYo: e.noYo})
	}
	return out
}
