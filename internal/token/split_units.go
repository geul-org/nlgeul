//ff:func feature=token type=reader control=iteration dimension=1
//ff:what splitUnits — runs with a run of digits and a Latin unit (20ms, 5kg) split in two: the number and the unit
package token

func splitUnits(s []rune, runs [][2]int) [][2]int {
	var out [][2]int
	for _, r := range runs {
		if cut, ok := unitCut(s, r); ok {
			out = append(out, [2]int{r[0], cut}, [2]int{cut, r[1]})
			continue
		}
		out = append(out, r)
	}
	return out
}
