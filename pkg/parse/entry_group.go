//ff:func feature=parse type=reader control=sequence
//ff:what entryGroup — the key readGroup files an entry's candidate under: groupBase (its lead, its form, its place or kind, the letters of a closing written wrong) and its name; candidates of different names are never filed together
package parse

func entryGroup(e *entry) string {
	return groupBase(e) + "·" + e.name
}
