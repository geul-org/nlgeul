//ff:func feature=token type=reader control=sequence
//ff:what entriesAt — the index entries that may end at j: those under the token before j, and under its keys as a fused token
package token

// The slice is the index's own when the index has worked it out (addEnds):
// it is read, never changed.

func (sc *scan) entriesAt(j int) []*entry {
	if sc.ix.ends != nil {
		return sc.ix.ends[sc.s[j-1]]
	}
	return sc.ix.entriesOf(sc.s[j-1])
}
