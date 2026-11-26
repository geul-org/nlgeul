//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what misfitsIn — how many of the entries are particles read against the norm (misfit) with the name given ("" for any name)
package parse

func misfitsIn(list []*entry, name string) int {
	n := 0
	for _, e := range list {
		if e.misfit && (name == "" || e.name == name) {
			n++
		}
	}
	return n
}
