//ff:func feature=parse type=reader control=sequence
//ff:what shortOf — one shortened closing read from a line of its table: the closing's row as the closing tables write one, the form it is short for, and the stems it follows; the mark that the copula does not take it is kept as the row's own
package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
)

func shortOf(row data.Row) (short, error) {
	m, err := MarkerOf(row, 0)
	if err != nil {
		return short{}, err
	}
	tail, err := row.Pick(5, tailNames)
	if err != nil {
		return short{}, err
	}
	bars := m.NoIda
	m.NoIda = false
	return short{m, row.Cells[4], bars, tail}, nil
}
