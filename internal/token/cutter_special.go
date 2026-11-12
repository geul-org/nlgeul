//ff:func feature=token type=reader control=sequence
//ff:what special — whether place i stands inside a stretch that is no sentence's word (an address, a mail address, a path, an option): the letters from where the unit's letters begin up to the next space are asked
package token

import "unicode"

func (c *cutter) special(i int) bool {
	from := min(c.at, i)
	end := from
	for end < len(c.r) && !unicode.IsSpace(c.r[end]) {
		end++
	}
	_, stop := special(c.r[from:end])
	return stop > 0 && i < from+stop
}
