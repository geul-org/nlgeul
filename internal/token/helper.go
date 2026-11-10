//ff:func feature=token type=reader control=sequence
//ff:what helper — the syllable a tail's stem ends in, when the tail attaches right to it: the syllable before the tail, as it was before a piece was fused into it; 0 when nothing stands there or the tail begins with the copula
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) helper(t tail) rune {
	if t.start < 1 || len(t.pieces) == 0 || t.pieces[0].tail&rule.TailCopula != 0 {
		return 0
	}
	if t.base != 0 {
		return t.base
	}
	return sc.s[t.start-1]
}
