//ff:func feature=mini type=helper control=sequence
//ff:what ctxFor — a rule context for one word and one of its tails
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func ctxFor(w *word, t tailFact) toulmin.Context {
	ctx := toulmin.NewContext()
	ctx.Set(keyWord, w)
	ctx.Set(keyTail, t)
	return ctx
}
