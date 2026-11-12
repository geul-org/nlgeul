//ff:func feature=token type=rule control=sequence
//ff:what stemVerdict — the verdict of a kind's graph for a stem with the facts given (what follows it, its observed count, its letters, the word before it); the same facts are judged once and remembered
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The key holds every fact the graph is given, so that a rule added to
// the graph is never answered from another stem's verdict.
func (ix *index) stemVerdict(kind int, p rule.Piece) float64 {
	key := stemKey{kind, p}
	if v, ok := ix.seen[key]; ok {
		return v
	}
	v := verdict(ix.kinds[kind], p)
	ix.seen[key] = v
	return v
}
