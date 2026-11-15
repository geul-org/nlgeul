//ff:func feature=data type=reader control=iteration dimension=1
//ff:what TestRowCount — a column holding a whole number of 0 or more gives it; a negative number and letters are errors that name the column
package data

import (
	"strings"
	"testing"
)

func TestRowCount(t *testing.T) {
	row := Row{File: "t.tsv", Line: 2, Names: []string{"횟수"}, Cells: []string{"12"}}
	if v, err := row.Count(0); err != nil || v != 12 {
		t.Fatal(v, err)
	}
	for _, bad := range []string{"-1", "열둘"} {
		row.Cells[0] = bad
		if _, err := row.Count(0); err == nil || !strings.Contains(err.Error(), "1번째 칸(횟수)") || !strings.Contains(err.Error(), "0 이상의 수가 아니다") {
			t.Errorf("%s: %v", bad, err)
		}
	}
}
