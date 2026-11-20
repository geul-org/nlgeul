//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what unitWordAt — does the stem that begins at b begin right after a number, its letters being a unit of the unit table and then a word written after a unit (afterUnitWords: 5년간 is 년 and 간)
package parse

import (
	"slices"
	"strings"
)

// The fact is read here; what it rebuts is the kinds' graphs' to judge.
func (sc *scan) unitWordAt(b int, stem string) bool {
	if b == 0 || !sc.numEnds[b] {
		return false
	}
	for _, u := range unitWords {
		if rest, ok := strings.CutPrefix(stem, u); ok && slices.Contains(afterUnitWords, rest) {
			return true
		}
	}
	return false
}
