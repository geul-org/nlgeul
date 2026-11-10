//ff:func feature=token type=reader control=iteration dimension=1
//ff:what splitUnits — runs with a run of digits and a Latin unit (20ms, 5kg) split in two: the number and the unit
package token

import "strings"

func splitUnits(s []rune, runs [][2]int) [][2]int {
	var out [][2]int
	for _, r := range runs {
		text := string(s[r[0]:r[1]])
		unit := strings.TrimLeft(text, "0123456789")
		if cut := r[1] - len([]rune(unit)); cut > r[0] && isLatinUnit(unit) {
			out = append(out, [2]int{r[0], cut}, [2]int{cut, r[1]})
			continue
		}
		out = append(out, r)
	}
	return out
}
