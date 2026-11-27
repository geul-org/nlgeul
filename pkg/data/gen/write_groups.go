//ff:func feature=data type=helper control=sequence
//ff:what writeGroups — writes the group default table: the three header lines and then the table lines, block by block
package main

import "io"

func writeGroups(w io.Writer, t *ucdTable, blocks []blockRange) error {
	return writeBlocks(w, tableHead("글자, 갈래, 이름, 값, 수 값", "groups"), t, blocks, blockGroupRows)
}
