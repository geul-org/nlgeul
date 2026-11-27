//ff:func feature=data type=helper control=iteration dimension=1
//ff:what deriveByName — the row of an enclosed character read from its name; found is whether a form key of the name stands in an enclosed block, ok whether the row could be made
package main

import (
	"slices"
	"strings"
)

func deriveByName(t *ucdTable, r rune, block string) (row formRow, found bool, ok bool) {
	head := strings.Split(t.Chars[r].Name, " WITH ")[0]
	for _, nf := range nameForms {
		if strings.Contains(head, nf[0]) && (slices.Contains(enclosedBlocks, block) || block == "Dingbats") {
			row, ok = nameFormRow(t, r, nf[1])
			return row, true, ok
		}
	}
	return formRow{}, false, false
}
