//ff:func feature=parse type=helper control=sequence
//ff:what headRun — the runs of a unit, the first begun one letter earlier when it follows a sign kept at the unit's head ($HOME, .env) and is no number
package parse

import (
	"strings"
	"unicode"
)

// runHeads are the signs kept at a unit's head that the first run takes
// in ($HOME, .env).
const runHeads = "$."

func (sc *scan) headRun(runs [][2]int) [][2]int {
	if len(runs) == 0 || runs[0][0] != 1 || !strings.ContainsRune(runHeads, sc.s[0]) || unicode.IsDigit(sc.s[1]) {
		return runs
	}
	return append([][2]int{{0, runs[0][1]}}, runs[1:]...)
}
