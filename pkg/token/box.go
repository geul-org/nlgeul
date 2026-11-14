//ff:type feature=token type=model
//ff:what box — what one node is drawn as before its width is fixed: the title, the rows inside (a row for each candidate shown and the row that counts those left out), and the width inside in cells that fits them all
package token

type box struct {
	title string
	rows  []string
	width int
}
