//ff:func feature=data type=helper control=sequence
//ff:what TestRowOpt — a column holding - has no value (empty), any other column is the letters written
package data

import "testing"

func TestRowOpt(t *testing.T) {
	row := Row{Cells: []string{"-", "값"}}
	if row.Opt(0) != "" || row.Opt(1) != "값" {
		t.Fatal(row.Opt(0), row.Opt(1))
	}
}
