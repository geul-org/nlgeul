//ff:func feature=data type=helper control=sequence
//ff:what TestLoadUCD — a file is read into a table and a missing file is an error
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadUCD(t *testing.T) {
	p := filepath.Join(t.TempDir(), "UnicodeData.txt")
	if err := os.WriteFile(p, []byte(sampleUCD), 0o644); err != nil {
		t.Fatal(err)
	}
	tbl, err := loadUCD(p)
	if err != nil || tbl.Chars[0x2460].Name != "CIRCLED DIGIT ONE" {
		t.Fatal(err)
	}
	if _, err := loadUCD(p + ".none"); err == nil {
		t.Fatal("missing")
	}
}
