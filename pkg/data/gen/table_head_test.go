//ff:func feature=data type=helper control=sequence
//ff:what TestTableHead — the head has the columns, the command with its mode, and the source line
package main

import "testing"

func TestTableHead(t *testing.T) {
	assertEq(t, tableHead("가, 나", "groups"), "# 칸(탭으로 나눔): 가, 나\n"+
		"# 기계로 만든 표다. 손으로 고치지 않는다: go run ./pkg/data/gen groups <UnicodeData.txt> <Blocks.txt> 로 다시 만든다.\n"+
		"# 자료: Unicode Character Database 15.0.0 (Unicode License v3, https://www.unicode.org/license.txt)\n")
}
