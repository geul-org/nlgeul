//ff:func feature=data type=helper control=sequence
//ff:what Row.Errf — an error about one column of a row: the file, the line, the column's number and name, what is written there, and the reason
package data

import "fmt"

func (r Row) Errf(col int, format string, a ...any) error {
	return fmt.Errorf("표 %s %d번째 줄 %d번째 칸(%s) \"%s\": %s", r.File, r.Line, col+1, r.Names[col], r.Cells[col], fmt.Sprintf(format, a...))
}
