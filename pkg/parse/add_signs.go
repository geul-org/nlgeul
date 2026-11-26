//ff:func feature=parse type=rule control=iteration dimension=2
//ff:what addSigns — a graph for each candidate of each named sign (signNames), by the sign (signGraphOf: the hyphen and the box-drawing line have graphs of their own); the full stop's are its closing candidates (마침표, 줄임표), and a middle dot (a sign named 가운뎃점) gets one more, 줄임표 for a run of it at an edge of its unit (dotsGraph: ······); a sign with no name that is read apart between letters (bareApart: =) gets a graph with no name to print (apartGraph)
package parse

import "slices"

func (ix *index) addSigns() {
	ix.marks = map[rune][]*entry{'.': ix.stops}
	dots, edges := dotsGraph()
	ellipsis := &entry{name: ellipsisName, g: dots}
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
	for _, sign := range signOrderNamed {
		for _, name := range signNames[sign] {
			g, edges := signGraphOf(sign, name)
			ix.marks[sign] = append(ix.marks[sign], &entry{name: name, g: g})
			ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
		}
		if slices.Contains(signNames[sign], middleDot) {
			ix.marks[sign] = append(ix.marks[sign], ellipsis)
		}
	}
	for _, sign := range bareApart {
		g, edges := apartGraph(sign)
		ix.marks[sign] = append(ix.marks[sign], &entry{g: g})
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
	}
}
