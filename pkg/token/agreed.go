//ff:func feature=token type=reader control=iteration dimension=2
//ff:what agreed — the boundaries every surviving candidate puts: the places where the unit is cut and sent up (section 12: cut by agreement, not by a threshold)
package token

import "sort"

func agreed(paths []*path) []int {
	count := map[int]int{}
	for _, p := range paths {
		for _, e := range wordsOf(p) {
			count[e.end]++
		}
	}
	var out []int
	for end, c := range count {
		if c == len(paths) {
			out = append(out, end)
		}
	}
	sort.Ints(out)
	return out
}
