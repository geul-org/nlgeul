//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wideRuns — the runs that splitUnits cuts in two (20ms), from where each begins to where it ends: the stretch that read as one stem is still a run of Latin letters and digits
package token

func wideRuns(s []rune, runs [][2]int) map[int]int {
	out := map[int]int{}
	for _, r := range runs {
		if _, ok := unitCut(s, r); ok {
			out[r[0]] = r[1]
		}
	}
	return out
}
