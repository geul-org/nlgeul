//ff:func feature=data type=reader control=sequence
//ff:what TestRowFlags — names of a closed set joined with | are summed as bits and - is none; a name outside the set and a name written twice are errors that name the column
package data

import (
	"strings"
	"testing"
)

func TestRowFlags(t *testing.T) {
	set := map[string]int{"가": 1, "나": 2, "다": 4}
	row := Row{File: "t.tsv", Line: 2, Names: []string{"갈래"}, Cells: []string{"가|다"}}
	if v, err := row.Flags(0, set); err != nil || v != 5 {
		t.Fatal(v, err)
	}
	row.Cells[0] = "-"
	if v, err := row.Flags(0, set); err != nil || v != 0 {
		t.Fatal(v, err)
	}
	row.Cells[0] = "가|라"
	if _, err := row.Flags(0, set); err == nil || !strings.Contains(err.Error(), "1번째 칸(갈래)") || !strings.Contains(err.Error(), "\"라\"은 정해진 이름이 아니다") || !strings.Contains(err.Error(), "가, 나, 다") {
		t.Fatal(err)
	}
	row.Cells[0] = "나|나"
	if _, err := row.Flags(0, set); err == nil || !strings.Contains(err.Error(), "1번째 칸(갈래)") || !strings.Contains(err.Error(), "\"나\"이 두 번 적혀 있다") {
		t.Fatal(err)
	}
}
