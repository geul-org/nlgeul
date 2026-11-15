//ff:func feature=token type=rule control=sequence
//ff:what addLoneShut — put the graph of a quotation mark that closes a pair with none of its kind open (loneShutGraph) in the index after the sign's other graphs, as a closing candidate, and count it
package token

func (ix *index) addLoneShut(sign rune, row pairRow) {
	g, edges := loneShutGraph(sign, row)
	ix.pairs[sign] = append(ix.pairs[sign], &entry{name: shutName + ":" + string(sign), form: shutName, g: g})
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
}
