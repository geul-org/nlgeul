//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what drawn — whether a stretch of one sign, from start to stop in some letters, stands in a run of signs that holds a sign that is no punctuation mark of a sentence (textMarks) and not the sign itself: an arrow's head (-->), a table's bar (|---|), another box-drawing letter (├──┼)
package parse

import "strings"

func drawn(all []rune, start, stop int) bool {
	from, to := start, stop
	for from > 0 && signRune(all[from-1]) {
		from--
	}
	for to < len(all) && signRune(all[to]) {
		to++
	}
	for k := from; k < to; k++ {
		if (k < start || k >= stop) && all[k] != all[start] && !strings.ContainsRune(textMarks, all[k]) {
			return true
		}
	}
	return false
}
