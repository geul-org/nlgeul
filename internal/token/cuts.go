//ff:func feature=token type=reader control=iteration dimension=2
//ff:what Cuts — for measuring: one line for each unit, its text and the word ends (token offsets) of each surviving reading, best first, readings separated by ";"
package token

import (
	"strconv"
	"strings"
)

func Cuts(sentence string) []string {
	ix := newIndex()
	var out []string
	for _, field := range strings.Fields(sentence) {
		sc := &scan{ix: ix, s: core(field)}
		var readings []string
		for _, p := range sc.search() {
			var ends []string
			for _, e := range wordsOf(p) {
				ends = append(ends, strconv.Itoa(e.end))
			}
			readings = append(readings, strings.Join(ends, ","))
		}
		out = append(out, string(sc.s)+"\t"+strings.Join(readings, ";"))
	}
	return out
}
