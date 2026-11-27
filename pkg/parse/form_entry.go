//ff:func feature=parse type=rule control=sequence
//ff:what formEntry — the entry of one name a sign of the tables of special letters is read by (원:로마자, 단위:제곱미터), made once and shared by every letter of that name: its graph is formGraph's: the claim "this letter is that", with nothing against it but for the kinds judged by the sign's place (가림, 글머리)
package parse

func (ix *index) formEntry(name string, drawn bool) *entry {
	if e, ok := ix.formEntries[name]; ok {
		return e
	}
	g, edges := formGraph(name, drawn)
	e := &entry{name: name, g: g}
	ix.formEntries[name], ix.graphs, ix.edges = e, ix.graphs+1, ix.edges+edges
	return e
}
