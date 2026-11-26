//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what cutMarked — whether the Hangul after the slash a unit's path is cut at reads as markers written onto the path (~/a/b/에, ~/a/b/이란, ~/a/b/인): some run of markers that ends at the unit's end starts at the cut, with nothing fused into the path and no stem of its own; false when the unit has no cut
package parse

// Such Hangul is a particle or an ending on a directory's slash before
// it is a file's name or a word of its own: the name reading is lowered
// (namedTail) and the path is not cut from it as from a word (wholeAt).
func (sc *scan) cutMarked() bool {
	if sc.cut == 0 {
		return false
	}
	for _, t := range sc.tailsAt(len(sc.s)) {
		if t.start == sc.cut && t.base == 0 && t.only == "" && t.as == "" && t.shown == "" && len(t.pieces) > 0 {
			return true
		}
	}
	return false
}
