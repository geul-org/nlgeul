//ff:func feature=data type=helper control=sequence
//ff:what TestWriteGroups — three header lines come first, then each block's lines, and a failing writer gives an error
package main

import (
	"strings"
	"testing"
)

func TestWriteGroups(t *testing.T) {
	tbl := loadSample(t)
	blocks := []blockRange{{Lo: 0x2460, Hi: 0x2488, Name: "Enclosed Alphanumerics"}, {Lo: 0x20, Hi: 0x7F, Name: "Basic Latin"}}
	var b strings.Builder
	if err := writeGroups(&b, tbl, blocks); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(b.String(), "\n"), "\n")
	if len(lines) != 4 || !strings.HasPrefix(lines[0], "# 칸(탭으로 나눔): 글자, 갈래") || lines[3] != ".\t기호\t그 밖\t.\t-" {
		t.Fatal(lines)
	}
	if writeGroups(failWriter{}, tbl, blocks) == nil {
		t.Fatal("head")
	}
	if writeGroups(&failAfterWriter{}, tbl, blocks) == nil {
		t.Fatal("rows")
	}
}
