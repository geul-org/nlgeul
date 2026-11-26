//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what quoted — whether a marker that takes a quotation (quoteMarkers: 라고, 고, 는, 라는, 란 …) is written right after the sign at place i: the Hangul syllables after the sign are such a marker, and no letter or digit stands after them ("라고 말했다, 다'는)
package parse

import (
	"slices"
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The marker is told by its letters alone, whatever it is read as.
func (c *cutter) quoted(i int) bool {
	j := i + 1
	for j < len(c.r) && rule.JongOf(c.r[j]) >= 0 {
		j++
	}
	next := atOf(c.r, j)
	return !unicode.IsLetter(next) && !unicode.IsDigit(next) && slices.Contains(quoteMarkers, string(c.r[i+1:j]))
}
