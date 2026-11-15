//ff:func feature=data type=reader control=iteration dimension=1
//ff:what Map — a table of two columns as a map from the first to the second; a first column written on two lines is an error
package data

func Map(name string) (map[string]string, error) {
	rows, err := Rows(name, 2)
	if err != nil {
		return nil, err
	}
	out := map[string]string{}
	for _, row := range rows {
		if _, ok := out[row.Cells[0]]; ok {
			return nil, row.Errf(0, "앞 줄에 이미 적혀 있다")
		}
		out[row.Cells[0]] = row.Cells[1]
	}
	return out, nil
}
