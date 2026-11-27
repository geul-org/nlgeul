//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what tailsAt — every run of markers that ends at j: particle chains, predicate tails, predicate tails with a closing 요, predicate tails with the particle 는 put under their closing as ㄴ (underClosings) or with 요 fused into the 지 their closing ends in (underYo), and any of these (or none) closed by a full stop; and an approximate number written in Hangul with no unit on it (approxLone: 육 and 칠 of 육칠)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) tailsAt(j int) []tail {
	if stops, start := sc.stopAt(j); len(stops) > 0 {
		var out []tail
		for _, stop := range stops {
			out = append(out, joined(sc.tailsAt(start), stop)...)
			out = append(out, tail{start: start, pieces: []piece{stop}})
		}
		return out
	}
	out := append(sc.chains(j, 0, -1), sc.closings(j, false, false)...)
	out = append(out, sc.fixed(j, -1)...)
	out = append(out, sc.units(j, 0, true)...)
	out = append(out, sc.tildeAt(j)...)
	out = append(out, sc.approxLone(j)...)
	out = append(out, sc.underClosings(j)...)
	out = append(out, sc.underYo(j)...)
	if j < 2 || sc.s[j-1] != '요' {
		return out
	}
	yo := piece{"요", yoName, 1, sc.ix.pref(sc.ix.named('요', yoName).count, "요"), rule.TailPred, false, ""}
	out = append(out, joined(sc.loneHonors(j-1, true), yo)...)
	for _, t := range sc.closings(j-1, true, false) {
		if !t.noYo {
			out = append(out, joined([]tail{t}, yo)...)
		}
	}
	return out
}
