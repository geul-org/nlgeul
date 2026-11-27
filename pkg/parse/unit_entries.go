//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what unitEntries — the entries of the letters given read as a unit of the unit table with no number before them (㎡ → 단위:제곱미터), one for each line of the table with those letters, in the order of the lines (formEntry); none when the table has no such letters
package parse

func (ix *index) unitEntries(form string) []*entry {
	var out []*entry
	for _, row := range unitRows {
		if row.Cells[0] == form {
			out = append(out, ix.formEntry(unitLabel+row.Cells[1], false))
		}
	}
	return out
}
