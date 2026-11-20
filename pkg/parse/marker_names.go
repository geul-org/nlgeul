//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what markerNames — the names of the marker graphs in the index (every entry under its keys, which markerGraph stamped)
package parse

func (ix *index) markerNames() map[string]bool {
	out := map[string]bool{}
	for _, list := range ix.byKey {
		for _, e := range list {
			out[e.name] = true
		}
	}
	return out
}
