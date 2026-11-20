//ff:func feature=parse type=helper control=iteration dimension=2
//ff:what entryNamed — the entry of the index with the name, under whatever key, or nil
package parse

import "slices"

func entryNamed(ix *index, name string) *entry {
	for _, list := range ix.byKey {
		if i := slices.IndexFunc(list, func(e *entry) bool { return e.name == name }); i >= 0 {
			return list[i]
		}
	}
	return nil
}
