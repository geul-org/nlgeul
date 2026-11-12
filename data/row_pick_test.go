//ff:func feature=data type=reader control=sequence
//ff:what TestRowPick — a column naming one of a closed set gives the value of the name; a name outside the set is an error that names the column and lists the set in the order of its letters
package data

import (
	"strings"
	"testing"
)

func TestRowPick(t *testing.T) {
	set := map[string]int{"둘": 2, "하나": 1}
	row := Row{File: "t.tsv", Line: 2, Names: []string{"수", "이름"}, Cells: []string{"하나", "셋"}}
	if v, err := row.Pick(0, set); err != nil || v != 1 {
		t.Fatal(v, err)
	}
	_, err := row.Pick(1, set)
	if err == nil || !strings.Contains(err.Error(), "2번째 칸(이름)") || !strings.Contains(err.Error(), "쓸 수 있는 이름: 둘, 하나") {
		t.Fatal(err)
	}
}
