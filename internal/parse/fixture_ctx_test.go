//ff:func feature=parse type=helper control=sequence
//ff:what ctxFor — a rule context for one word and one of its tails
package parse

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func ctxFor(w *word, t tailFact) toulmin.Context {
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyWord, facts(w))
	ctx.Set(rule.KeyTail, tailFacts(t))
	return ctx
}
