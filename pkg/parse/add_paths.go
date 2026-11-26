//ff:func feature=parse type=rule control=sequence
//ff:what addPaths — the graphs of the two readings of a path that the letters alone do not settle, with their verdicts worked out: a path read with the Hangul after its last slash as its name (pathTailName: its verdict, and the one when those letters read as markers alone), and a path with no head and one slash whose name holds Hangul joined by a sign (pathWeakName)
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// Neither reading is taken away: a rebuttal lowers it and the other
// readings of the same letters stay beside it. Hangul right after a
// directory's slash that reads as markers alone is a particle written
// onto the path (~/a/b/에) before it is a file's name. One slash and no
// head are weak grounds for a path: the same letters are two words set
// side by side (TV/라디오).
func (ix *index) addPaths() {
	tail := toulmin.NewGraph(pathTailName)
	claim := tail.Rule(rule.PieceRead).With(&rule.Claim{Name: pathTailName, Concludes: "기호"})
	tail.Counter(rule.PieceIs).With(fact(rule.PieceMarked)).With(&rule.Claim{Name: whyPathMarked, Concludes: "반박"}).Attacks(claim, forcePathMarked)
	ix.pathNamed, ix.pathMarked = verdict(tail, rule.Piece{}), verdict(tail, rule.Piece{Marked: true})
	weak := toulmin.NewGraph(pathWeakName)
	stands := weak.Rule(rule.PieceRead).With(&rule.Claim{Name: pathWeakName, Concludes: "기호"})
	weak.Counter(rule.PieceRead).With(&rule.Claim{Name: whyPathWeak, Concludes: "반박"}).Attacks(stands, forcePathWeak)
	ix.pathWeak = verdict(weak, rule.Piece{})
	ix.graphs, ix.edges = ix.graphs+2, ix.edges+2
}
