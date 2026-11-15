//ff:func feature=data type=reader control=iteration dimension=1
//ff:what Column — the first column of a table of the columns given, line by line, each value once, in the order it is first written
package data

import "slices"

func Column(name string, cols int) ([]string, error) {
	rows, err := Rows(name, cols)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, row := range rows {
		if !slices.Contains(out, row.Cells[0]) {
			out = append(out, row.Cells[0])
		}
	}
	return out, nil
}
