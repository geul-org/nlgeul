//ff:func feature=parse type=reader control=sequence
//ff:what sides — the facts of the sign at place i that tell where a quotation mark is written: a marker that takes a quotation stands right after it (MarkQuoted, quoted); a space or the sentence's start stands right before it and a letter or digit right after it, and no such marker (MarkOpenSide: an opening mark's place); it is MarkQuoted, or a letter stands right before it and a space or the sentence's end right after it (MarkShutSide: a closing mark's place)
package parse

import (
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// An opening quotation mark is written onto the word after it and a
// closing one onto the word before it (한글 맞춤법, 부록 문장 부호), and
// the markers that take a quotation are written onto its closing mark.
func (c *cutter) sides(i int) int {
	left, right := atOf(c.r, i-1), atOf(c.r, i+1)
	apartLeft := left == 0 || unicode.IsSpace(left)
	apartRight := right == 0 || unicode.IsSpace(right)
	facts := 0
	if c.quoted(i) {
		facts |= rule.MarkQuoted | rule.MarkShutSide
	}
	if apartLeft && facts == 0 && (unicode.IsLetter(right) || unicode.IsDigit(right)) {
		facts |= rule.MarkOpenSide
	}
	if !apartLeft && sideOf(string(left)) != 0 && apartRight {
		facts |= rule.MarkShutSide
	}
	return facts
}
