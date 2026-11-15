//ff:func feature=data type=reader control=iteration dimension=1
//ff:what Each — a table of the columns given as a list of values, one made from each line by the function given, in the order of the lines; the first line that makes an error ends the reading
package data

func Each[T any](name string, cols int, of func(Row) (T, error)) ([]T, error) {
	rows, err := Rows(name, cols)
	if err != nil {
		return nil, err
	}
	out := make([]T, 0, len(rows))
	for _, row := range rows {
		v, err := of(row)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}
