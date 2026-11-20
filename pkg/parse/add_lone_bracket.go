//ff:func feature=parse type=rule control=sequence
//ff:what addLoneBracket — put the graph of a closing bracket that closes a pair with none of its kind open (loneBracketGraph) in the index after the sign's other graph, as a closing candidate, and count it
package parse

func (ix *index) addLoneBracket(sign rune, row pairRow) {
	g, edges := loneBracketGraph(sign, row)
	ix.pairs[sign] = append(ix.pairs[sign], &entry{name: shutName + ":" + string(sign), form: shutName, g: g})
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
}
