//ff:func feature=token type=reader control=iteration dimension=2
//ff:what endsOf — for measuring: the word ends (token offsets) of each candidate, joined with ",", the candidates joined with ";"
package token

import (
	"strconv"
	"strings"
)

func endsOf(paths []*path) string {
	var readings []string
	for _, p := range paths {
		var ends []string
		for _, e := range wordsOf(p) {
			ends = append(ends, strconv.Itoa(e.end))
		}
		readings = append(readings, strings.Join(ends, ","))
	}
	return strings.Join(readings, ";")
}
