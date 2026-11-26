//ff:func feature=parse type=reader control=sequence
//ff:what jointAt — whether a joint that hands a number on to the next stands at a place of the unit: one of caseJoints (의, 제) and a digit right after it (의4 in 31조의4, 제1 in 36조제1항)
package parse

import "slices"

func (sc *scan) jointAt(k int) bool {
	return slices.Contains(caseJoints, string(sc.at(k))) && digitsFrom(sc.s, k+1) > 0
}
