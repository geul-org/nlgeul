//ff:type feature=parse type=model
//ff:what sentenceRule — a rule that judges a candidate of a word with the words of its sentence around it: its name (as a trace prints it), its graph (its claim about the candidate there, and the exceptions against it), the preference a candidate gains when the claim stands (bonus), its force when it is a rule of rebuttal (force: the claim is that the grammar does not let the candidate stand there, and when it stands the candidate's verdict is multiplied by 2/(1+force)-1; 0 for a rule of preference), which candidates it speaks of at all (about), and which of them cannot be judged before the words after them are read up to the predicate (waits; nil when it reads nothing after the candidate)
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A rule of a sentence never takes a candidate away. A rule of
// preference (force 0) adds to the preference of those its claim stands
// for and never lowers one. A rule of rebuttal (force above 0, the one
// value forceUngrammatical) lowers the verdict of those its claim stands
// for as a rebuttal of the unit layer does, and leaves their preference
// as it is; when more than one stands against a candidate, their
// factors are multiplied. The candidates of the word are put in order by
// lambda again.
type sentenceRule struct {
	name  string
	graph *toulmin.Graph
	bonus float64
	force float64
	about func(rule.Near) bool
	waits func(rule.Near) bool
}
