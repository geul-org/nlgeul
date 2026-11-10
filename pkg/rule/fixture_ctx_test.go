//ff:func feature=rule type=helper control=sequence
//ff:what ctxFor — a rule context for one word and one tail reading
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func ctxFor(w *Word, t Tail) toulmin.Context {
	ctx := toulmin.NewContext()
	ctx.Set(KeyWord, w)
	ctx.Set(KeyTail, t)
	return ctx
}
