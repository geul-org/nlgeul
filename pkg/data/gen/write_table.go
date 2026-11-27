//ff:func feature=data type=helper control=iteration dimension=1
//ff:what writeTable — writes the three header lines and then the table lines, block by block
package main

import (
	"fmt"
	"io"
	"strings"
)

func writeTable(w io.Writer, t *ucdTable, blocks []blockRange) error {
	head := "# 칸(탭으로 나눔): 글자, 꼴, 내용, 값, 수 값\n" +
		"# 기계로 만든 표다. 손으로 고치지 않는다: go run ./pkg/data/gen <UnicodeData.txt> <Blocks.txt> 로 다시 만든다.\n" +
		"# 자료: Unicode Character Database " + unicodeVersion + " (Unicode License v3, https://www.unicode.org/license.txt)\n"
	if _, err := io.WriteString(w, head); err != nil {
		return err
	}
	jamo := jamoMap(t)
	for _, b := range blocks {
		rows := blockRows(t, jamo, b)
		if len(rows) == 0 {
			continue
		}
		if _, err := fmt.Fprintln(w, strings.Join(rows, "\n")); err != nil {
			return err
		}
	}
	return nil
}
