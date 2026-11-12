//ff:func feature=token type=reader control=sequence
//ff:what keptAt — the readings that end at q where a sign between letters stands: those kept at q, and those that end with a word of no markers there (a sign shows the boundary, so none is held back for its place, whole with shown); worked out once for each place
package token

func (sc *scan) keptAt(best [][]*path, q int) []*path {
	if got, ok := sc.kept[q]; ok {
		return got
	}
	got := sc.ix.top(sc.whole(append([]*path{}, best[q]...), best, q, true))
	sc.kept[q] = got
	return got
}
