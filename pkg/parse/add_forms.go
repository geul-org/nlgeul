//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what addForms — the graphs of the letters of the tables of special letters that are signs (signRune: ⓐ, ㉠, ㈎, ？), read at a unit's edge as the sign contest of that letter: one graph for each name (shape:content) of the tables, shared by the letters of that name, the claim "this letter is that shape of that content" with nothing against it (any number of the letter in a row); a letter of several lines has a candidate for each; a letter the sign table names (ix.marks: ，, ＋) is left to that table; and the verdict of a letter of the tables that is a letter of a run of Latin letters and digits read alone by its shape (loneForm, the graph loneFormGraph: the claim and against it "the letter is a letter of a run", with the force a new word is weighed with)
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) addForms() {
	alone := toulmin.NewGraph(loneFormGraph)
	shape := alone.Rule(rule.PieceRead).With(&rule.Claim{Name: loneFormGraph, Concludes: "열쇠"})
	alone.Counter(rule.PieceRead).With(&rule.Claim{Name: whyRunLetter, Concludes: "반박"}).Attacks(shape, forceNew)
	ix.loneForm, ix.graphs, ix.edges = verdict(alone, rule.Piece{}), ix.graphs+1, ix.edges+1
	ix.forms = map[rune][]*entry{}
	byName := map[string]*entry{}
	for _, f := range formList {
		if !signRune(f.letter) || len(ix.marks[f.letter]) > 0 {
			continue
		}
		e, ok := byName[f.name()]
		if !ok {
			g, edges := signGraph(f.name())
			e = &entry{name: f.name(), g: g}
			byName[f.name()], ix.graphs, ix.edges = e, ix.graphs+1, ix.edges+edges
		}
		ix.forms[f.letter] = append(ix.forms[f.letter], e)
	}
}
