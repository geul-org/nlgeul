//ff:func feature=data type=helper control=sequence
//ff:what Row.Opt — a column that may have no value: "" where - is written, else the letters
package data

func (r Row) Opt(col int) string {
	if r.Cells[col] == none {
		return ""
	}
	return r.Cells[col]
}
