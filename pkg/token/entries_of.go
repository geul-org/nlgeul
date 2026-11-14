//ff:func feature=token type=reader control=iteration dimension=1
//ff:what entriesOf — the index entries that may end with the token r: those under the token, and under its keys as a fused token (keysAt), in that order
package token

func (ix *index) entriesOf(r rune) []*entry {
	var out []*entry
	for _, k := range keysAt(r) {
		out = append(out, ix.byKey[k]...)
	}
	return out
}
