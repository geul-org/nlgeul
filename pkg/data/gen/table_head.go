//ff:func feature=data type=helper control=sequence
//ff:what tableHead — the three header lines of a table: its columns, the command that makes it, and the source of the data
package main

func tableHead(columns, mode string) string {
	return "# 칸(탭으로 나눔): " + columns + "\n" +
		"# 기계로 만든 표다. 손으로 고치지 않는다: go run ./pkg/data/gen " + mode + " <UnicodeData.txt> <Blocks.txt> 로 다시 만든다.\n" +
		"# 자료: Unicode Character Database " + unicodeVersion + " (Unicode License v3, https://www.unicode.org/license.txt)\n"
}
