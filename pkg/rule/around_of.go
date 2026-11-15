//ff:func feature=rule type=helper control=sequence
//ff:what AroundOf — the candidate under judgment with the words around it (an empty one when the context has none)
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func AroundOf(ctx toulmin.Context) Around {
	v, _ := ctx.Get(KeyAround)
	a, _ := v.(Around)
	return a
}
