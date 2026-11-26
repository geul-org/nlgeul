//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what lineBare — whether the input line that holds place i of a sentence holds no letter and no digit: a line of signs alone (---, |---|---|, -----------)
package parse

import "unicode"

func (c *cutter) lineBare(i int) bool {
	from, to := i, i
	for from > 0 && c.r[from-1] != '\n' {
		from--
	}
	for to < len(c.r) && c.r[to] != '\n' {
		to++
	}
	for _, r := range c.r[from:to] {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
