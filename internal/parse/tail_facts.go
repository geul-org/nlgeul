//ff:func feature=parse type=reader control=sequence
//ff:what tailFacts — one tail reading as the rule package's facts
package parse

import "github.com/geul-org/koparse/pkg/rule"

func tailFacts(t tailFact) rule.Tail {
	return rule.Tail{Closing: t.closing, Kind: t.kind, Stem: t.stem, Size: t.size, Copula: t.copula, Count: t.count, Marked: t.marked}
}
