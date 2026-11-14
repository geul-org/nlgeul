//ff:func feature=token type=reader control=sequence
//ff:what facts — the facts of the sign at place i of a sentence that do not hang on what stands open: what stands right on each side, whether a space or the sentence's edge stands on both sides, and whether a hyphen or an equals sign stands right next to it
package token

import (
	"strings"
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (c *cutter) facts(i int) rule.Mark {
	left, right := atOf(c.r, i-1), atOf(c.r, i+1)
	m := rule.Mark{Left: sideOf(string(left)), Right: sideOf(string(right)), Run: 1}
	if left == 0 {
		m.Left = 0
	}
	if right == 0 {
		m.Right = 0
	}
	if (left == 0 || unicode.IsSpace(left)) && (right == 0 || unicode.IsSpace(right)) {
		m.Facts |= rule.MarkApart
	}
	if (left != 0 && strings.ContainsRune(arrowSigns, left)) || (right != 0 && strings.ContainsRune(arrowSigns, right)) {
		m.Facts |= rule.MarkArrow
	}
	return m
}
