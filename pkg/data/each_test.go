//ff:func feature=data type=reader control=sequence
//ff:what TestEach — a value is made from each line by the function given, in the order of the lines; a table that is not there is an error, and the first line the function refuses ends the reading with its error
package data

import (
	"strings"
	"testing"
)

func TestEach(t *testing.T) {
	useTables(t, map[string]string{"ok": head + "이름, 횟수\n가\t5\n나\t7\n", "bad": head + "이름, 횟수\n가\t5\n나\t많이\n다\t9\n"})
	count := func(r Row) (int, error) { return r.Count(1) }
	got, err := Each("ok", 2, count)
	if err != nil || len(got) != 2 || got[0] != 5 || got[1] != 7 {
		t.Fatal(got, err)
	}
	if _, err := Each("none", 2, count); err == nil || !strings.Contains(err.Error(), "표 none.tsv") {
		t.Fatal(err)
	}
	if _, err := Each("bad", 2, count); err == nil || !strings.Contains(err.Error(), "표 bad.tsv 3번째 줄 2번째 칸(횟수)") {
		t.Fatal(err)
	}
}
