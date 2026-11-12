//ff:func feature=token type=reader control=sequence
//ff:what TestLoadWords — the table of testdata is read with comment lines skipped and a key's readings folded into one; a missing file and a bad line are errors that say which
package token

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadWords(t *testing.T) {
	lex, typos, total, err := loadWords("testdata/words.tsv")
	if err != nil || lex["작"].kinds != kindNominal|kindAdj || lex["작"].count != 7000 || lex["작"].by[2] != 7000 || lex["작"].by[0] != 400 || lex["고양이"].count != 1000 || total <= 7000 || len(typos) != 0 {
		t.Fatal(lex["작"], total, err)
	}
	if _, _, _, err := loadWords("testdata/none.tsv"); err == nil || !strings.Contains(err.Error(), "낱말 표를 열 수 없다") {
		t.Fatal(err)
	}
	bad := filepath.Join(t.TempDir(), "bad.tsv")
	os.WriteFile(bad, []byte("# x\n\n고양이\t명사\t\t10\n고양이 명사\n"), 0o600)
	if _, _, _, err := loadWords(bad); err == nil || !strings.Contains(err.Error(), "4번째 줄") {
		t.Fatal(err)
	}
}
