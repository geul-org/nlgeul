//ff:func feature=data type=reader control=sequence
//ff:what TestColumn — the first column is given line by line with each value once, in the order it is first written; a table that is not there is an error
package data

import (
	"strings"
	"testing"
)

func TestColumn(t *testing.T) {
	useTables(t, map[string]string{"ok": head + "꼴, 뜻\n나\t1\n가\t2\n나\t3\n"})
	got, err := Column("ok", 2)
	if err != nil || strings.Join(got, ",") != "나,가" {
		t.Fatal(got, err)
	}
	if _, err := Column("none", 2); err == nil || !strings.Contains(err.Error(), "표 none.tsv") {
		t.Fatal(err)
	}
}
