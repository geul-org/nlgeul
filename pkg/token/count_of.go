//ff:func feature=token type=reader control=sequence
//ff:what countOf — the observed count of a key as one kind (by the kind's place in kindOrder): that of its rows of the kind, 0 for a row with no count, and -1 when the kind is the key's only from a closed set of the grammar tables, which give no count
package token

func countOf(w word, kind, place int) int {
	if w.kinds&kind == 0 {
		return -1
	}
	return w.by[place]
}
