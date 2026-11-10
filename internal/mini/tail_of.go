//ff:func feature=mini type=helper control=sequence
//ff:what tailOf — the tail reading under judgment (an unread tail, syll 0, when the word has none)
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func tailOf(ctx toulmin.Context) tailFact {
	v, _ := ctx.Get(keyTail)
	t, _ := v.(tailFact)
	return t
}
