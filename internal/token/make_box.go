//ff:func feature=token type=helper control=sequence
//ff:what makeBox — the box of a node: its title and rows (boxRows with head), and the width inside that holds the widest row, or the title with two dashes at least
package token

func makeBox(n Node, head int) box {
	b := box{title: boxTitle(n), rows: boxRows(n, head)}
	b.width = max(maxCells(b.rows), cellWidth(b.title)+2)
	return b
}
