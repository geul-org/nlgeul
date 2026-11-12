//ff:func feature=token type=rule control=sequence
//ff:what addPair — put the graph of one candidate of a sign of a pair in the index, and count it
package token

func (ix *index) addPair(sign rune, row pairRow, name string) {
	g, edges := pairGraph(sign, row, name)
	ix.pairs[sign] = append(ix.pairs[sign], &entry{name: name + ":" + string(sign), form: name, g: g})
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
}
