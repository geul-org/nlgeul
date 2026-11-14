//ff:func feature=token type=reader control=iteration dimension=1
//ff:what addEnds — work out once, for every token an entry may end with (each key of the index and each Hangul syllable), the entries that may end with it (entriesOf), so that a scan does not put them together at every place
package token

// It is called when the index is whole; add drops what was worked out.
func (ix *index) addEnds() {
	ends := map[rune][]*entry{}
	for k := range ix.byKey {
		ends[k] = ix.entriesOf(k)
	}
	for r := rune(0xAC00); r <= 0xD7A3; r++ {
		if es := ix.entriesOf(r); len(es) > 0 {
			ends[r] = es
		}
	}
	ix.ends = ends
}
