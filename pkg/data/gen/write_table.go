//ff:func feature=data type=helper control=sequence
//ff:what writeTable — writes the form table: the three header lines and then the table lines, block by block
package main

import "io"

func writeTable(w io.Writer, t *ucdTable, blocks []blockRange) error {
	return writeBlocks(w, tableHead("글자, 꼴, 내용, 값, 수 값", "forms"), t, blocks, blockRows)
}
