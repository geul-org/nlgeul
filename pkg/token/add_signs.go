//ff:func feature=token type=rule control=iteration dimension=2
//ff:what addSigns — a graph for each candidate of each named sign (signNames), by the sign; the full stop's are its closing candidates (마침표, 줄임표)
package token

func (ix *index) addSigns() {
	ix.marks = map[rune][]*entry{'.': ix.stops}
	for _, sign := range signOrderNamed {
		for _, name := range signNames[sign] {
			g, edges := signGraph(name)
			ix.marks[sign] = append(ix.marks[sign], &entry{name: name, g: g})
			ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
		}
	}
}
