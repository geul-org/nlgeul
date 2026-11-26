//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what misfitOf — the first of the entries that is a particle read against the norm (misfit) with the name given; an empty entry when there is none
package parse

func misfitOf(list []*entry, name string) *entry {
	for _, e := range list {
		if e.misfit && e.name == name {
			return e
		}
	}
	return &entry{}
}
