//ff:func feature=parse type=reader control=sequence
//ff:what endingOf — one row of the closing table read from a line: the closing as MarkerOf reads it from the first four columns, and its observed count from the fifth
package parse

import "github.com/geul-org/nlgeul/data"

func endingOf(row data.Row) (ending, error) {
	m, err := MarkerOf(row, 0)
	if err != nil {
		return ending{}, err
	}
	count, err := row.Count(4)
	if err != nil {
		return ending{}, err
	}
	flags, _ := row.Flags(3, FlagNames)
	return ending{m.Form, m.Lead, m.After, m.Kind, flags, count}, nil
}
