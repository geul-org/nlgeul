//ff:func feature=parse type=rule control=sequence
//ff:what formEntry — the entry of one name a sign of the tables of special letters is read by (원:로마자, 단위:제곱미터), made once and shared by every letter of that name: its graph is the claim "this letter is that" with nothing against it (signGraph of a name it has no exception for)
package parse

func (ix *index) formEntry(name string) *entry {
	if e, ok := ix.formEntries[name]; ok {
		return e
	}
	g, edges := signGraph(name)
	e := &entry{name: name, g: g}
	ix.formEntries[name], ix.graphs, ix.edges = e, ix.graphs+1, ix.edges+edges
	return e
}
