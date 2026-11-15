//ff:func feature=data type=reader control=sequence
//ff:what TestRows — a table is read by its name from where the package reads its tables, with its file name tsv on every row; one that is not there is an error that names the file, and one written wrong is the error rowsOf gives
package data

import (
	"strings"
	"testing"
)

func TestRows(t *testing.T) {
	useTables(t, map[string]string{"ok": head + "꼴, 수\n가\t1\n", "bad": head + "꼴, 수\n가\n"})
	rows, err := Rows("ok", 2)
	if err != nil || len(rows) != 1 || rows[0].File != "ok.tsv" || rows[0].Line != 2 {
		t.Fatal(rows, err)
	}
	if _, err := Rows("none", 2); err == nil || !strings.Contains(err.Error(), "표 none.tsv") {
		t.Fatal(err)
	}
	if _, err := Rows("bad", 2); err == nil || !strings.Contains(err.Error(), "표 bad.tsv 2번째 줄") {
		t.Fatal(err)
	}
}
