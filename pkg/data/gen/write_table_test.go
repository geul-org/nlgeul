//ff:func feature=data type=helper control=sequence
//ff:what TestWriteTable — three header lines come first, then each block's lines, a block with none is skipped, and a failing writer gives an error
package main

import (
	"strings"
	"testing"
)

func TestWriteTable(t *testing.T) {
	tbl := loadSample(t)
	blocks := []blockRange{{Lo: 0x2460, Hi: 0x2488, Name: "Enclosed Alphanumerics"}, {Lo: 0x41, Hi: 0x5A, Name: "Basic Latin"}}
	var b strings.Builder
	if err := writeTable(&b, tbl, blocks); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(b.String(), "\n"), "\n")
	if len(lines) != 6 || !strings.HasPrefix(lines[2], "# 자료: Unicode Character Database 15.0.0") || !strings.HasPrefix(lines[3], "①\t") {
		t.Fatal(lines)
	}
	if writeTable(failWriter{}, tbl, blocks) == nil {
		t.Fatal("head")
	}
}
