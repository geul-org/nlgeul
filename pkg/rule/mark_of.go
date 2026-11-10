//ff:func feature=rule type=helper control=sequence
//ff:what MarkOf — the sign under judgment (an empty one when the context has none)
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func MarkOf(ctx toulmin.Context) Mark {
	v, _ := ctx.Get(KeyMark)
	m, _ := v.(Mark)
	return m
}
