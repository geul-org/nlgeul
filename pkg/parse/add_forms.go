//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what addForms — the graphs of the letters of the tables of special letters that are signs (signRune: ⓐ, ㉠, ㈎, ？), read at a unit's edge as the sign contest of that letter: one graph for each name (shape:content) of the tables, shared by the letters of that name, the claim "this letter is that shape of that content" with nothing against it (any number of the letter in a row); a letter of several lines has a candidate for each; a letter the sign table names (ix.marks: ，, ＋) is left to that table
package parse

func (ix *index) addForms() {
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
