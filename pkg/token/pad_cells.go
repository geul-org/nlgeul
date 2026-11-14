//ff:func feature=token type=helper control=sequence
//ff:what padCells — the text with spaces after it up to the width in cells (as it is when it is wider)
package token

import "strings"

func padCells(s string, width int) string {
	return s + strings.Repeat(" ", max(0, width-cellWidth(s)))
}
