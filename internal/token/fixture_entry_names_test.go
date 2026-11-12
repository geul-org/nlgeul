//ff:func feature=token type=helper control=iteration dimension=1
//ff:what entryNames — the names of some index entries, in order
package token

func entryNames(es []*entry) []string {
	var out []string
	for _, e := range es {
		out = append(out, e.name)
	}
	return out
}
