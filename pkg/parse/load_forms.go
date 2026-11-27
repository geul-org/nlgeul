//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what loadForms — the lines of the tables of special letters, in the order of the tables given and of their lines: a letter, its shape, what it writes, the letters it writes and the number it stands for (- for none), a line a form; a letter may have lines of more than one name
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/data"
)

func loadForms(tables ...[]data.Row) ([]form, error) {
	var out []form
	for _, row := range slices.Concat(tables...) {
		letter, err := row.Rune(0)
		if err != nil {
			return nil, err
		}
		out = append(out, form{letter, row.Cells[1], row.Cells[2], row.Cells[3], row.Opt(4)})
	}
	return out, nil
}
