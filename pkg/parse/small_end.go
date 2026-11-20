//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what smallEnd — where the longest run of Sino-Korean digits and 십, 백, 천 that begins at place i ends (i itself when there is none)
package parse

import "strings"

func (sc *scan) smallEnd(i int) int {
	for i < len(sc.s) && strings.ContainsRune(sinoDigits+"십백천", sc.s[i]) {
		i++
	}
	return i
}
