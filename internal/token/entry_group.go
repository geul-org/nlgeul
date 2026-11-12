//ff:func feature=token type=reader control=sequence
//ff:what entryGroup — the key readGroup files an entry's reading under: groupBase (its lead, its form, its place or kind, the letters of a closing written wrong) and its name; readings of different names are never filed together
package token

func entryGroup(e *entry) string {
	return groupBase(e) + "·" + e.name
}
