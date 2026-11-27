//ff:func feature=data type=helper control=sequence
//ff:what TestLoadBlocks — a file is read into blocks and a version, and a missing file is an error
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadBlocks(t *testing.T) {
	p := filepath.Join(t.TempDir(), "Blocks.txt")
	if err := os.WriteFile(p, []byte("# Blocks-15.0.0.txt\n0000..007F; Basic Latin\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	blocks, version, err := loadBlocks(p)
	if err != nil || version != "15.0.0" || len(blocks) != 1 {
		t.Fatal(blocks, version, err)
	}
	if _, _, err := loadBlocks(p + ".none"); err == nil {
		t.Fatal("missing")
	}
}
