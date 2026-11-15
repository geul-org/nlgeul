//ff:func feature=data type=helper control=sequence
//ff:what TestRowErrf — an error about a column gives the file, the line, the column's number and name, what is written there and the reason made from the format
package data

import "testing"

func TestRowErrf(t *testing.T) {
	row := Row{File: "t.tsv", Line: 3, Names: []string{"꼴", "수"}, Cells: []string{"가", "여섯"}}
	want := `표 t.tsv 3번째 줄 2번째 칸(수) "여섯": 뜻 5`
	if got := row.Errf(1, "뜻 %d", 5).Error(); got != want {
		t.Fatal(got)
	}
}
