//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what namedEntries — the entries among those given that have a name to print (a candidate with no name, the equals sign read apart, is left out)
package parse

func namedEntries(list []*entry) []*entry {
	var out []*entry
	for _, e := range list {
		if e.name != "" {
			out = append(out, e)
		}
	}
	return out
}
