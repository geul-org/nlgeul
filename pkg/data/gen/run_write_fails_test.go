//ff:func feature=data type=helper control=iteration dimension=1
//ff:what TestRunWriteFails — a table longer than the write buffer reaches the writer before the flush, and a failing writer then stops run with an error
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunWriteFails(t *testing.T) {
	dir := t.TempDir()
	var u strings.Builder
	for i := 0; i < 0xA0; i++ {
		fmt.Fprintf(&u, "%04X;CIRCLED DIGIT ONE;No;0;ON;<circle> 0031;;1;1;N;;;;;\n", 0x2460+i)
	}
	blocks := "# Blocks-15.0.0.txt\n" + strings.Repeat("2460..24FF; Enclosed Alphanumerics\n", 8)
	ucd, bl := filepath.Join(dir, "UnicodeData.txt"), filepath.Join(dir, "Blocks.txt")
	for p, s := range map[string]string{ucd: u.String(), bl: blocks} {
		if err := os.WriteFile(p, []byte(s), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if run("forms", ucd, bl, failWriter{}) == nil {
		t.Fatal("the failing writer must stop run")
	}
	var b strings.Builder
	if err := run("forms", ucd, bl, &b); err != nil || b.Len() <= 4096 {
		t.Fatal(err, b.Len())
	}
}
