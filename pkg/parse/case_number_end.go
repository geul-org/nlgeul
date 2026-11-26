//ff:func feature=parse type=reader control=sequence
//ff:what caseNumberEnd — where the shape of a court case number ends that begins at a place of the unit (0 when none begins there): a year (caseYearLens digits), a mark of 1 to caseMarkMax Hangul syllables and a serial number (a run of digits that no Latin letter runs on from); not when the mark is a greater number (17만2000) or ends with a joint between two numbers (caseJointMark: 11의2, 31조의4, 3000분의7, 36조제1항), nor when it is a unit of measure there (caseUnitMark: 11초92, 2020년3월, 36도5분)
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) caseNumberEnd(from int) int {
	year := digitsFrom(sc.s, from)
	mark := syllablesFrom(sc.s, from+year)
	at := from + year + mark
	end := at + digitsFrom(sc.s, at)
	if !slices.Contains(caseYearLens, year) || mark < 1 || mark > caseMarkMax || end == at || rule.RunRune(sc.at(end)) {
		return 0
	}
	marked := string(sc.s[from+year : at])
	if sc.caseJointMark(marked) || sc.readsAfter(marked, placeBig) || sc.caseUnitMark(marked, end) {
		return 0
	}
	return end
}
