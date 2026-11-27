//ff:func feature=data type=helper control=sequence
//ff:what TestWriteBlocks — a head comes first and then the lines of each block, a block with no line adds nothing, and a failing writer gives an error at the head and at a block
package main

import (
	"strings"
	"testing"
)

func TestWriteBlocks(t *testing.T) {
	tbl := loadSample(t)
	blocks := []blockRange{{Lo: 1, Hi: 1, Name: "A"}, {Lo: 2, Hi: 2, Name: "B"}, {Lo: 3, Hi: 3, Name: "C"}}
	rows := func(_ *ucdTable, _ map[rune]rune, b blockRange) []string {
		if b.Name == "B" {
			return nil
		}
		return []string{b.Name + "1", b.Name + "2"}
	}
	var b strings.Builder
	if err := writeBlocks(&b, "head\n", tbl, blocks, rows); err != nil {
		t.Fatal(err)
	}
	assertEq(t, b.String(), "head\nA1\nA2\nC1\nC2\n")
	if writeBlocks(failWriter{}, "head\n", tbl, blocks, rows) == nil {
		t.Fatal("head")
	}
	if writeBlocks(&failAfterWriter{}, "head\n", tbl, blocks, rows) == nil {
		t.Fatal("block")
	}
}
