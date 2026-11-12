//ff:func feature=token type=reader control=iteration dimension=1
//ff:what afterSuffixes — a copula tail read once more after each particle chain whose last piece is a suffix of copulaSuffixes and that ends where the tail starts (꽃+들+이+다): the tail keeps what it tells of its closing, and a dropped stem keeps the gate of the word before the suffix
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) afterSuffixes(t tail) []tail {
	if t.base != 0 {
		return nil
	}
	var out []tail
	for _, l := range sc.chains(t.start, 1, rule.LinkCopula) {
		if !slices.Contains(copulaSuffixes, l.pieces[len(l.pieces)-1].names) {
			continue
		}
		l = joined([]tail{l}, t.pieces...)[0]
		l.noYo, l.close = t.noYo, t.close
		if l.gate == nil && t.gate != nil {
			l.gate = t.gate
		}
		out = append(out, l)
	}
	return out
}
