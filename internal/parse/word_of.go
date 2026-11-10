//ff:func feature=parse type=helper control=sequence
//ff:what wordOf — the facts of the word under judgment (an empty word in an empty sentence when the context has none)
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func wordOf(ctx toulmin.Context) *word {
	v, _ := ctx.Get(keyWord)
	if w, ok := v.(*word); ok {
		return w
	}
	return &word{sent: &sentence{}}
}
