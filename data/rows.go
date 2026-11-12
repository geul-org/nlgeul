//ff:func feature=data type=reader control=sequence
//ff:what Rows — the lines of one table held in the program (tables/<name>.tsv), read and checked as rowsOf does
package data

import (
	"fmt"
	"io/fs"
)

func Rows(name string, cols int) ([]Row, error) {
	file := name + ".tsv"
	text, err := fs.ReadFile(tables, "tables/"+file)
	if err != nil {
		return nil, fmt.Errorf("표 %s: 품은 표에 없다", file)
	}
	return rowsOf(file, string(text), cols)
}
