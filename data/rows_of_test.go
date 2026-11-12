//ff:func feature=data type=reader control=iteration dimension=2
//ff:what TestRowsOf — a table's text is read into rows with the file, the line number counted over skipped lines too, the column names and the columns; a first line that is not the names, a line with another number of columns, an empty column and a line written twice are errors that say the file, the line and the reason
package data

import (
	"strings"
	"testing"
)

func TestRowsOf(t *testing.T) {
	text := head + "꼴, 수\n# 설명\n\n가\t1\n나\t2"
	rows, err := rowsOf("a.tsv", text, 2)
	if err != nil || len(rows) != 2 {
		t.Fatal(rows, err)
	}
	if rows[0].File != "a.tsv" || rows[0].Line != 4 || rows[1].Line != 5 || rows[1].Cells[1] != "2" || rows[0].Names[1] != "수" {
		t.Fatal(rows)
	}
	cases := []struct {
		name, text string
		want       []string
	}{
		{"no names", "꼴, 수\n가\t1\n", []string{"표 a.tsv 1번째 줄", "칸 이름 2개"}},
		{"too few names", head + "꼴\n가\t1\n", []string{"표 a.tsv 1번째 줄", "칸 이름 2개"}},
		{"empty", "", []string{"표 a.tsv 1번째 줄"}},
		{"cells", head + "꼴, 수\n가\n", []string{"표 a.tsv 2번째 줄", "칸이 1개다", "꼴, 수"}},
		{"twice", head + "꼴, 수\n가\t1\n나\t2\n가\t1\n", []string{"표 a.tsv 4번째 줄", "2번째 줄과 같은 줄"}},
		{"empty cell", head + "꼴, 수\n가\t\n", []string{"표 a.tsv 2번째 줄 2번째 칸(수)", "빈 칸"}},
	}
	for _, c := range cases {
		_, err := rowsOf("a.tsv", c.text, 2)
		if err == nil {
			t.Errorf("%s: no error", c.name)
			continue
		}
		for _, w := range c.want {
			if !strings.Contains(err.Error(), w) {
				t.Errorf("%s: %q lacks %q", c.name, err, w)
			}
		}
	}
}
