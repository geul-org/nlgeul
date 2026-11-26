//ff:func feature=parse type=reader control=sequence
//ff:what caseUnitMark — whether the letters between the two numbers of a case number's shape are a unit of measure there and no case mark: a word of the unit table (11초92, 10시30, 81승53: a number, its unit and a number), but a unit that is a case mark too (caseUnitMarks: 97도2836) or another piece read after a number only when such a piece is read right after the second number (36도5분, 2020년3월: a measure written with no space)
package parse

import "slices"

func (sc *scan) caseUnitMark(mark string, end int) bool {
	if sc.readsAfter(mark, placeUnit) && !slices.Contains(caseUnitMarks, mark) {
		return true
	}
	return sc.readsAfter(mark, 0) && sc.afterAt(end)
}
