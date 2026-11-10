//ff:func feature=rule type=helper control=sequence
//ff:what markCtx — a rule context for one sign
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func markCtx(m Mark) toulmin.Context {
	ctx := toulmin.NewContext()
	ctx.Set(KeyMark, m)
	return ctx
}
