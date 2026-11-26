//ff:func feature=parse type=rule control=sequence
//ff:what addSwap — put the graph of a round quotation mark read against its shape (swapGraph) in the index after the sign's other graphs, as an opening candidate of a mark drawn as a closing one and a closing candidate of a mark drawn as an opening one, and count it
package parse

func (ix *index) addSwap(sign rune) {
	g, name, edges := swapGraph(sign)
	ix.pairs[sign] = append(ix.pairs[sign], &entry{name: name + ":" + string(sign), form: name, g: g})
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
}
