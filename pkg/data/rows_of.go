//ff:func feature=data type=reader control=iteration dimension=1
//ff:what rowsOf — read a table's text: the first line names the columns ("# 칸(탭으로 나눔): " and the names with ", " between them), a line that begins with "# " or is empty is skipped, every other line is a row of the columns the caller asks for, tab-separated; a line with another number of columns, an empty column or the same letters as a line before it is an error that names the file, the line and the reason
package data

import (
	"fmt"
	"strings"
)

func rowsOf(file, text string, cols int) ([]Row, error) {
	lines := strings.Split(strings.TrimSuffix(text, "\n"), "\n")
	names := strings.Split(strings.TrimPrefix(lines[0], head), ", ")
	if !strings.HasPrefix(lines[0], head) || len(names) != cols {
		return nil, fmt.Errorf("표 %s 1번째 줄: 첫 줄은 \"%s\" 뒤에 칸 이름 %d개를 쉼표와 빈칸으로 나누어 적어야 한다", file, head, cols)
	}
	var out []Row
	seen := map[string]int{}
	for i, line := range lines {
		if strings.HasPrefix(line, "# ") || line == "" {
			continue
		}
		row := Row{File: file, Line: i + 1, Names: names, Cells: strings.Split(line, "\t")}
		if err := row.check(seen[line]); err != nil {
			return nil, err
		}
		seen[line] = i + 1
		out = append(out, row)
	}
	return out, nil
}
