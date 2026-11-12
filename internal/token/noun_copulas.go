//ff:func feature=token type=reader control=iteration dimension=1
//ff:what nounCopulas — the candidates of some predicate tails after a dependent noun the copula's stem is fused into (테니, 텐데: 터+(이)+니, 터+(이)+(ㄴ)데): those of each tail (nounCopula)
package token

func (sc *scan) nounCopulas(tails []tail) []tail {
	var out []tail
	for _, t := range tails {
		out = append(out, sc.nounCopula(t)...)
	}
	return out
}
