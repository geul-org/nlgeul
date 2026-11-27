//ff:func feature=data type=helper control=iteration dimension=1
//ff:what TestRun — a forms or groups table is written from two files, an unknown mode is an error, a wrong Blocks.txt version stops it, and a missing file is an error
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	dir := t.TempDir()
	ucd := filepath.Join(dir, "UnicodeData.txt")
	ok := filepath.Join(dir, "Blocks.txt")
	old := filepath.Join(dir, "OldBlocks.txt")
	files := map[string]string{ucd: sampleUCD,
		ok:  "# Blocks-15.0.0.txt\n2460..24FF; Enclosed Alphanumerics\n",
		old: "# Blocks-14.0.0.txt\n2460..24FF; Enclosed Alphanumerics\n"}
	for p, s := range files {
		if err := os.WriteFile(p, []byte(s), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	var b strings.Builder
	if err := run("forms", ucd, ok, &b); err != nil || !strings.Contains(b.String(), "①\t원\t수\t1\t1\n") {
		t.Fatal(err, b.String())
	}
	if err := run("forms", ucd, old, &b); err == nil || !strings.Contains(err.Error(), "14.0.0") {
		t.Fatal(err)
	}
	b.Reset()
	if err := run("groups", ucd, ok, &b); err != nil || !strings.Contains(b.String(), "글자, 갈래, 이름") {
		t.Fatal(err, b.String())
	}
	if err := run("none", ucd, ok, &b); err == nil || !strings.Contains(err.Error(), "none") {
		t.Fatal(err)
	}
	if run("forms", ucd+".none", ok, &b) == nil || run("forms", ucd, ok+".none", &b) == nil {
		t.Fatal("missing")
	}
}
