//ff:func feature=rule type=helper control=sequence
//ff:what WordOf — the facts of the word under judgment (an empty word in an empty sentence when the context has none)
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func WordOf(ctx toulmin.Context) *Word {
	v, _ := ctx.Get(KeyWord)
	if w, ok := v.(*Word); ok {
		return w
	}
	return &Word{Sent: &Sentence{}}
}
