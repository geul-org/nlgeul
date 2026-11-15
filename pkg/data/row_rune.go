//ff:func feature=data type=reader control=sequence
//ff:what Row.Rune — the value of a column that holds one letter
package data

import "unicode/utf8"

func (r Row) Rune(col int) (rune, error) {
	if utf8.RuneCountInString(r.Cells[col]) != 1 {
		return 0, r.Errf(col, "글자 하나가 아니다")
	}
	v, _ := utf8.DecodeRuneInString(r.Cells[col])
	return v, nil
}
