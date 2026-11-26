//ff:func feature=parse type=reader control=selection
//ff:what lean — the candidate a quotation mark written the same way at both ends (" ') leans to by what is written on its sides: an opening one has a space or the sentence's start on its left and a letter or digit on its right ('결혼); a closing one has a marker that takes a quotation right after it ("라고, '는: quoted), or something other than a space on its left and a space or the sentence's end on its right (된다', 다.'); "" for any other sign or place
package parse

import "unicode"

// An opening quotation mark is written onto the word after it and a
// closing one onto the word before it (한글 맞춤법, 부록 문장 부호: the
// marks are written close to what they quote).
func (c *cutter) lean(i int, row pairRow) string {
	left, right := atOf(c.r, i-1), atOf(c.r, i+1)
	apartLeft := left == 0 || unicode.IsSpace(left)
	apartRight := right == 0 || unicode.IsSpace(right)
	switch {
	case !row.opens || !row.shuts:
		return ""
	case c.quoted(i):
		return shutName
	case apartLeft && (unicode.IsLetter(right) || unicode.IsDigit(right)):
		return openName
	case !apartLeft && apartRight:
		return shutName
	}
	return ""
}
