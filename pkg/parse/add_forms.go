//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what addForms — the candidates of the letters of the tables of special letters that are signs (signRune: ⓐ, ㉠, ㈎, ？, ，, ㎡), read at a unit's edge after the candidates of the sign table (signContest), in the order they are printed: the units of the unit table written with that letter (㎡[단위:제곱미터]: the hand-written table first), then the names (kind:what, the table's 꼴 and 내용) the tables give the letter, one for each of its lines, then the readings through the letters it writes (valueEntries: ？[물음표] after ？[전각:기호]); the letters of one name share one graph, the claim with nothing against it (formEntry); and the verdict of a letter of the tables that is a letter of a run of Latin letters and digits read alone by its shape (loneForm, the graph loneFormGraph: the claim and against it "the letter is a letter of a run", with the force a new word is weighed with)
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
	ix.forms, ix.formEntries = map[rune][]*entry{}, map[string]*entry{}
	for _, f := range formList {
		if !signRune(f.letter) || (f.origin == defaultTable+".tsv" && ix.handNamed(f.letter)) {
			continue
		}
		if len(ix.forms[f.letter]) > 0 {
			ix.forms[f.letter] = append(ix.forms[f.letter], ix.formEntry(f.name(), f.origin != madeTable+".tsv"))
			continue
		}
		ix.forms[f.letter] = append(append(ix.unitEntries(string(f.letter)), ix.formEntry(f.name(), f.origin != madeTable+".tsv")), ix.valueEntries(f)...)
	}
}
