//ff:type feature=token type=model
//ff:what sentenceRule — a rule that judges a candidate of a word with the words of its sentence around it: its name (as a trace prints it), its graph (the claim that the candidate is preferred there, and the exceptions against it), the preference a candidate gains when the claim stands, which candidates it speaks of at all (about), and which of them cannot be judged before the words after them are read up to the predicate (waits; nil when it reads nothing after the candidate)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A rule of a sentence never takes a candidate away and never lowers
// one: it adds to the preference of those its claim stands for, and the
// candidates of the word are put in order by lambda again.
type sentenceRule struct {
	name  string
	graph *toulmin.Graph
	bonus float64
	about func(rule.Near) bool
	waits func(rule.Near) bool
}
