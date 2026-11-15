//ff:func feature=data type=reader control=iteration dimension=1
//ff:what Runes — the first column of a table whose first column is one letter, line by line, each letter once, in the order it is first written
package data

import "slices"

func Runes(name string, cols int) ([]rune, error) {
	rows, err := Rows(name, cols)
	if err != nil {
		return nil, err
	}
	var out []rune
	for _, row := range rows {
		r, err := row.Rune(0)
		if err != nil {
			return nil, err
		}
		if !slices.Contains(out, r) {
			out = append(out, r)
		}
	}
	return out, nil
}
