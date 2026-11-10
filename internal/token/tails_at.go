//ff:func feature=token type=reader control=iteration dimension=1
//ff:what tailsAt — every run of markers that ends at j: particle chains, predicate tails, and predicate tails with a closing 요
package token

func (sc *scan) tailsAt(j int) []tail {
	out := append(sc.chains(j, 0, -1), sc.closings(j, false)...)
	if j < 2 || sc.s[j-1] != '요' {
		return out
	}
	yo := piece{"요", "보조사:요", 1, sc.ix.pref(2894, "요")}
	for _, t := range sc.closings(j-1, true) {
		if !t.noYo {
			out = append(out, joined([]tail{t}, yo)...)
		}
	}
	return out
}
