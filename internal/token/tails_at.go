//ff:func feature=token type=reader control=iteration dimension=1
//ff:what tailsAt — every run of markers that ends at j: particle chains, predicate tails, predicate tails with a closing 요, and any of these (or none) closed by a full stop
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) tailsAt(j int) []tail {
	if stop, start, ok := sc.stopAt(j); ok {
		return append(joined(sc.tailsAt(start), stop), tail{start: start, pieces: []piece{stop}})
	}
	out := append(sc.chains(j, 0, -1), sc.closings(j, false)...)
	out = append(out, sc.fixed(j, -1)...)
	out = append(out, sc.units(j, 0, true)...)
	out = append(out, sc.tildeAt(j)...)
	if j < 2 || sc.s[j-1] != '요' {
		return out
	}
	yo := piece{"요", yoName, 1, sc.ix.pref(sc.ix.named('요', yoName).count, "요"), rule.TailPred, false}
	out = append(out, joined(sc.loneHonors(j-1, true), yo)...)
	for _, t := range sc.closings(j-1, true) {
		if !t.noYo {
			out = append(out, joined([]tail{t}, yo)...)
		}
	}
	return out
}
