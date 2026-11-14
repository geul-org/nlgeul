//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lines — the box drawn at the width inside given (the box's own width or a wider one): the title line "┌ title ───┐", a line "│ row │" for each row, and the bottom line "└───┘"; every line is width + 4 cells wide
package token

import "strings"

func (b box) lines(width int) []string {
	out := make([]string, 0, len(b.rows)+2)
	out = append(out, "┌ "+b.title+" "+strings.Repeat("─", width-cellWidth(b.title))+"┐")
	for _, row := range b.rows {
		out = append(out, "│ "+padCells(row, width)+" │")
	}
	return append(out, "└"+strings.Repeat("─", width+2)+"┘")
}
