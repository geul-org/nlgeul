//ff:func feature=token type=helper control=sequence
//ff:what region — the region of a stretch of a unit at a place of the sentence: made from its cells when the net read it (newRegionNet), else from the ways of cutting its branches make (newRegion); c is the stretch's contest with the unit and Glued filled in
package token

func (st unitStretch) region(c Contest, at int) *region {
	if st.cells != nil {
		return newRegionNet(c, at, *st.cells)
	}
	return newRegion(c, at, cutsOf(c, at))
}
