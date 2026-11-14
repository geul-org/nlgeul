//ff:func feature=token type=reader control=iteration dimension=2
//ff:what marksOf — for measuring: the word ends (token offsets) of each candidate joined with ",", and after each end that is inside the unit ":" and the names of the spacing rules that are on and that the boundary is a place of, joined with "/" ("-" when there is none); the candidates joined with ";"
package token

import (
	"strconv"
	"strings"
)

// The names come from spacingOf, judged on the two words that meet at
// the boundary and the letters each takes up.
func (sc *scan) marksOf(paths []*path) string {
	var candidates []string
	for _, p := range paths {
		words := wordsOf(p)
		var ends []string
		for i, w := range words {
			end := strconv.Itoa(w.end)
			if i+1 < len(words) {
				end += ":" + sc.spacingNames(words, i)
			}
			ends = append(ends, end)
		}
		candidates = append(candidates, strings.Join(ends, ","))
	}
	return strings.Join(candidates, ";")
}
