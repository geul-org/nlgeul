//ff:func feature=data type=reader control=iteration dimension=1
//ff:what Counts — a table of two columns, a name and a count, as a map; a name written on two lines is an error
package data

func Counts(name string) (map[string]int, error) {
	rows, err := Rows(name, 2)
	if err != nil {
		return nil, err
	}
	out := map[string]int{}
	for _, row := range rows {
		n, err := row.Count(1)
		if err != nil {
			return nil, err
		}
		if _, ok := out[row.Cells[0]]; ok {
			return nil, row.Errf(0, "앞 줄에 이미 적혀 있다")
		}
		out[row.Cells[0]] = n
	}
	return out, nil
}
