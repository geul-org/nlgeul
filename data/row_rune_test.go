//ff:func feature=data type=reader control=iteration dimension=1
//ff:what TestRowRune — a column holding one letter gives it, a letter of more than one byte too; two letters are an error that names the column
package data

import (
	"strings"
	"testing"
)

func TestRowRune(t *testing.T) {
	row := Row{File: "t.tsv", Line: 2, Names: []string{"기호"}, Cells: []string{"가"}}
	for _, ok := range []string{"가", "("} {
		row.Cells[0] = ok
		if v, err := row.Rune(0); err != nil || string(v) != ok {
			t.Errorf("%s: %q %v", ok, v, err)
		}
	}
	row.Cells[0] = "가나"
	if _, err := row.Rune(0); err == nil || !strings.Contains(err.Error(), "1번째 칸(기호)") || !strings.Contains(err.Error(), "글자 하나가 아니다") {
		t.Fatal(err)
	}
}
