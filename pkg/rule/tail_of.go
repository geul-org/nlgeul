//ff:func feature=rule type=helper control=sequence
//ff:what TailOf — the tail reading under judgment (an unread tail, Size 0, when the context has none)
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func TailOf(ctx toulmin.Context) Tail {
	v, _ := ctx.Get(KeyTail)
	t, _ := v.(Tail)
	return t
}
