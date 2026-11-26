//ff:func feature=parse type=reader control=selection
//ff:what namedTail — the path a stem from b to end is by one of the two readings that the letters alone do not settle: the whole unit as a path with the Hangul after its path's last slash as its name (sc.cut: ~/a/b/회의록), with that reading's verdict, lowered when those Hangul letters read as markers on the path (sc.onSlash, ~/a/b/에: a particle written onto a directory's slash); or the unit's letters up to where they end as a path with no head and one slash (sc.weak: notes/회의-결과), with that reading's verdict; nothing for any other stem
package parse

func (sc *scan) namedTail(b, end int) []markCandidate {
	switch {
	case b != 0:
		return nil
	case sc.cut > 0 && end == len(sc.s) && sc.onSlash:
		return []markCandidate{{pathName, sc.ix.pathMarked}}
	case sc.cut > 0 && end == len(sc.s):
		return []markCandidate{{pathName, sc.ix.pathNamed}}
	case sc.weak > 0 && end == sc.weak:
		return []markCandidate{{pathName, sc.ix.pathWeak}}
	}
	return nil
}
