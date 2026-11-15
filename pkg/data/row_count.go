//ff:func feature=data type=reader control=sequence
//ff:what Row.Count — the value of a column that holds a count: a whole number, 0 or more
package data

import "strconv"

func (r Row) Count(col int) (int, error) {
	n, err := strconv.Atoi(r.Cells[col])
	if err != nil || n < 0 {
		return 0, r.Errf(col, "0 이상의 수가 아니다")
	}
	return n, nil
}
