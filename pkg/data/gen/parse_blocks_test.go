//ff:func feature=data type=helper control=sequence
//ff:what TestParseBlocks — the version of the first line and the blocks are read, comment lines are skipped, a code that cannot be read is an error
package main

import (
	"strings"
	"testing"
)

func TestParseBlocks(t *testing.T) {
	in := "# Blocks-15.0.0.txt\n# Date: x\n\n0000..007F; Basic Latin\n2150..218F; Number Forms\n"
	blocks, version, err := parseBlocks(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}
	assertEq(t, version, "15.0.0")
	if len(blocks) != 2 || blocks[1].Lo != 0x2150 || blocks[1].Hi != 0x218F || blocks[1].Name != "Number Forms" {
		t.Fatal(blocks)
	}
	_, version, err = parseBlocks(strings.NewReader("0000..007F; A\n"))
	if err != nil || version != "" {
		t.Fatal(version, err)
	}
	if _, _, err := parseBlocks(strings.NewReader("0000..0FFFFFFFFFFF; Big\n")); err == nil {
		t.Fatal("too big")
	}
}
