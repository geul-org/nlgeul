//ff:func feature=data type=helper control=iteration dimension=1
//ff:what writeBlocks — writes a head and then the lines the rows function gives for each block, skipping a block with none
package main

import (
	"fmt"
	"io"
	"strings"
)

func writeBlocks(w io.Writer, head string, t *ucdTable, blocks []blockRange,
	rows func(*ucdTable, map[rune]rune, blockRange) []string) error {
	if _, err := io.WriteString(w, head); err != nil {
		return err
	}
	jamo := jamoMap(t)
	for _, b := range blocks {
		lines := rows(t, jamo, b)
		if len(lines) == 0 {
			continue
		}
		if _, err := fmt.Fprintln(w, strings.Join(lines, "\n")); err != nil {
			return err
		}
	}
	return nil
}
