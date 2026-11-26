//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what notTail — a closing's tail read once more for the stem 아니 alone: its pieces as a tail any predicate takes, so that the adjective 아니 stands before a closing that otherwise only verbs take
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func notTail(t tail) tail {
	ps := append([]piece{}, t.pieces...)
	for i := range ps {
		ps[i].tail = rule.TailPred
	}
	t.pieces, t.only = ps, notOnly
	return t
}
