//ff:func feature=data type=reader control=sequence
//ff:what Row.Pick — the value of a column that holds one of a closed set of names; a name that is not of the set is an error that lists the set
package data

func (r Row) Pick(col int, names map[string]int) (int, error) {
	v, ok := names[r.Cells[col]]
	if !ok {
		return 0, r.Errf(col, "정해진 이름이 아니다. 쓸 수 있는 이름: %s", listOf(names))
	}
	return v, nil
}
