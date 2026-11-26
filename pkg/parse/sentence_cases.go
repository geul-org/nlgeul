//ff:func feature=parse type=rule control=sequence
//ff:what sentenceCases — put the rows of the case exception table under a rule of a sentence's name on its claim as one exception: the candidate's word (its letters, rule.ListSelfText) is one of the rows' words (whyCase; a trace adds the row: caseNote); nothing (nil) when no row names the rule, else that exception, so a rule can be given a counter to it
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A row names a word the rule's claim does not hold for although its
// grammar does, found case by case, with the ground for it; the reading
// it gives is the one the rule would have taken away.
func sentenceCases(g *toulmin.Graph, claim *toulmin.Rule, name string) *toulmin.Rule {
	cases := caseExceptions[name]
	if len(cases) == 0 {
		return nil
	}
	c := except(g, rule.Listed, whyCase, listed("caseExceptions", rule.ListSelfText, caseWords(cases)), &caseRows{Rows: cases})
	c.Attacks(claim)
	return c
}
