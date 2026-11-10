//ff:func feature=token type=reader control=iteration dimension=1
//ff:what entriesAt — the index entries that may end at j: those under the token before j, and under its keys as a fused token
package token

func (sc *scan) entriesAt(j int) []*entry {
	var out []*entry
	for _, k := range keysAt(sc.s[j-1]) {
		out = append(out, sc.ix.byKey[k]...)
	}
	return out
}
