//ff:func feature=parse type=helper control=sequence
//ff:what rowFor — a row of a table made from its columns, for a test of the readers of rows: file t.tsv, line 2, and as many column names as columns
package parse

import "github.com/geul-org/nlgeul/data"

func rowFor(cells ...string) data.Row {
	names := []string{"꼴", "앞말", "갈래", "조건", "횟수", "여섯째", "일곱째"}
	return data.Row{File: "t.tsv", Line: 2, Names: names[:len(cells)], Cells: cells}
}
