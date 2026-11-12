//ff:func feature=token type=reader control=selection
//ff:what next — the unit that begins at place i of a sentence and where it ends: a span of code with what is written onto it; a sign that cuts the sentence (cutSigns) alone; a sign of a pair, alone when it only opens, with what is written onto it when it may close; else the letters up to the next space, cut sign, sign of a pair or span of code; the ways the signs may stand open are moved past a sign of a pair
package token

import "unicode"

func (c *cutter) next(i int) (unit, int) {
	u := unit{first: i == 0 || c.r[i-1] == '\n'}
	end := i + 1
	j, coded := c.code[i]
	reads, bare, open, notes := c.reads(i)
	cut, cutNotes := c.cuts(i)
	c.at = i
	switch {
	case coded:
		c.at = j + 1
		u.head, u.name, end = j+1-i, codeName, c.ontoEnd(j+1)
	case cut:
		u.notes = cutNotes
	case len(reads) > 0:
		u.head, u.name, u.reads, u.bare, u.notes = 1, pairName, reads, bare, notes
		c.at = i + 1
		end = c.afterSign(i+1, reads)
		c.open, u.measure = c.opened(i, open)
	default:
		end = c.wordEnd(i + 1)
	}
	u.text, u.spaced = string(c.r[i:end]), end < len(c.r) && unicode.IsSpace(c.r[end])
	return u, end
}
