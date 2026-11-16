//ff:func feature=token type=reader control=iteration dimension=1
//ff:what afterLinked — a copula tail read once more after each predicate tail that ends where the tail starts with a closing of copulaClosings (찾+아서+이+다, 위하+(여)서+(이)+다): the copula's pieces make a predicate tail there, for the stem before the closing is a predicate's, not a nominal's; the tail keeps what it tells of its closing and its gate
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) afterLinked(t tail) []tail {
	if t.base != 0 || t.start < 2 || !lastOfAny(copulaClosings, sc.s[t.start-1]) {
		return nil
	}
	rest := append([]piece{}, t.pieces...)
	for i := range rest {
		rest[i].tail = rule.TailPred
	}
	var out []tail
	for _, l := range sc.closings(t.start, false, false) {
		if !slices.Contains(copulaClosings, l.pieces[len(l.pieces)-1].names) {
			continue
		}
		l = joined([]tail{l}, rest...)[0]
		l.noYo, l.close = t.noYo, t.close
		if l.gate == nil && t.gate != nil {
			l.gate = t.gate
		}
		out = append(out, l)
	}
	return out
}
