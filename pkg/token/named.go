//ff:func feature=token type=reader control=iteration dimension=1
//ff:what named — the index entry under a key that has a name (not one that reads a closing written wrong under the name of the right one), or an entry with no count and no graph when there is none
package token

func (ix *index) named(key rune, name string) *entry {
	for _, e := range ix.byKey[key] {
		if e.name == name && e.wrote == "" {
			return e
		}
	}
	return &entry{name: name}
}
