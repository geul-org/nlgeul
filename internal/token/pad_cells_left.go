//ff:func feature=token type=helper control=sequence
//ff:what padCellsLeft — the text with spaces before it up to the width in cells (as it is when it is wider)
package token

import "strings"

func padCellsLeft(s string, width int) string {
	return strings.Repeat(" ", max(0, width-cellWidth(s))) + s
}
