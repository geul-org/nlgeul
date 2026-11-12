//ff:func feature=data type=reader control=sequence
//ff:what TestRowCheck — a row of the columns its table names passes; one with another number of columns, one written before at a line, and one with an empty column are each an error that names the file and the line
package data

import (
	"strings"
	"testing"
)

func TestRowCheck(t *testing.T) {
	row := Row{File: "t.tsv", Line: 7, Names: []string{"꼴", "수"}, Cells: []string{"가", "1"}}
	if err := row.check(0); err != nil {
		t.Fatal(err)
	}
	err := Row{File: "t.tsv", Line: 7, Names: row.Names, Cells: []string{"가"}}.check(0)
	if err == nil || !strings.Contains(err.Error(), "표 t.tsv 7번째 줄") || !strings.Contains(err.Error(), "칸이 1개다") || !strings.Contains(err.Error(), "꼴, 수") {
		t.Fatal(err)
	}
	err = row.check(3)
	if err == nil || !strings.Contains(err.Error(), "표 t.tsv 7번째 줄") || !strings.Contains(err.Error(), "3번째 줄과 같은 줄") {
		t.Fatal(err)
	}
	err = Row{File: "t.tsv", Line: 7, Names: row.Names, Cells: []string{"가", ""}}.check(0)
	if err == nil || !strings.Contains(err.Error(), "표 t.tsv 7번째 줄 2번째 칸(수)") || !strings.Contains(err.Error(), "빈 칸") {
		t.Fatal(err)
	}
}
