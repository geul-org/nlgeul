//ff:func feature=token type=reader control=iteration dimension=1
//ff:what addTypoClosings — the entry of every row of the table of closings written wrong, made from the entry of the right closing the row names (a row whose right closing the tables do not have makes none)
package token

func (ix *index) addTypoClosings(rights map[string]*entry) {
	for _, row := range typoClosings {
		if right := rights[row.Name]; right != nil {
			ix.addTypo(typoClosingEntry(right, row))
		}
	}
}
