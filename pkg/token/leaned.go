//ff:func feature=token type=reader control=iteration dimension=1
//ff:what leaned — the ways the signs before the sign at place i may stand open, in the order the sign is read in them: for a quotation mark that leans to a candidate by its sides (lean), the ways in which it reads as that candidate come first (for an opening one the ways with none of its kind open, for a closing one the ways with one open), each part in the order it was; else the ways as they are
package token

import "strings"

// The first way after the sign is the way of its first candidate, so
// the signs after it are read on from the candidate it leans to.
func (c *cutter) leaned(i int, row pairRow) []string {
	lean := c.lean(i, row)
	if lean == "" {
		return c.open
	}
	var first, rest []string
	for _, way := range c.open {
		if strings.ContainsRune(way, row.kind) == (lean == shutName) {
			first = append(first, way)
		} else {
			rest = append(rest, way)
		}
	}
	return append(first, rest...)
}
