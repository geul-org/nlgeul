//ff:func feature=token type=helper control=iteration dimension=1
//ff:what wordEnd — where the letters from a place of a sentence stop: at the next space, at the next sign that is read as a sign of a pair in some way the signs before it may stand open, at the next sign that cuts the sentence, at the next span of code, or at the end
package token

import "unicode"

func (c *cutter) wordEnd(from int) int {
	for ; from < len(c.r) && !unicode.IsSpace(c.r[from]); from++ {
		_, coded := c.code[from]
		candidates, _, _, _ := c.candidates(from)
		cut, _ := c.cuts(from)
		if coded || len(candidates) > 0 || cut {
			break
		}
	}
	return from
}
