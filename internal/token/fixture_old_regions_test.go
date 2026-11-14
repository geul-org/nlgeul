//ff:func feature=token type=helper control=iteration dimension=2
//ff:what oldRegions — the regions of a sentence as the stream made them before the net: each unit read whole by readUnit, each contest cut by its branches (newRegion with cutsOf), at the places the stream puts them
package token

func oldRegions(r *Reader, sentence string) []*region {
	var out []*region
	cursor, spaced := 0, true
	for i, u := range unitsOf(r.ix, sentence, false) {
		contests, _ := readUnit(r.ix, u, false)
		for k, c := range contests {
			c.Unit, c.Glued = i, k > 0 || (i > 0 && !spaced)
			out = append(out, newRegion(c, cursor, cutsOf(c, cursor)))
			cursor += len([]rune(c.Text))
		}
		spaced = u.spaced
	}
	return out
}
