//ff:func feature=token type=helper control=iteration dimension=1
//ff:what caseNote — what a trace adds to the reason of an exception that applied: for a case exception (one that carries caseRows) the word it matched and its row of the case exception table, for any other ""
package token

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// evidence is what the rule function gave back: for rule.Listed the
// fact it looked up, here the word under judgment.
func caseNote(specs toulmin.Specs, evidence any) string {
	for _, s := range specs {
		if c, ok := s.(*caseRows); ok {
			return c.note(evidence)
		}
	}
	return ""
}
