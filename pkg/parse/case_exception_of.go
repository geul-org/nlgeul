//ff:func feature=parse type=reader control=sequence
//ff:what caseExceptionOf — one row of the case exception table read from a line of its table: the graph, the word, the reading it keeps (- for none), what it does (지움 only), the kind of its ground (규범 or 관찰) and the ground; another name in those two columns is an error that lists the names
package parse

import "github.com/geul-org/nlgeul/pkg/data"

func caseExceptionOf(row data.Row) (caseException, error) {
	if _, err := row.Pick(3, caseStrengths); err != nil {
		return caseException{}, err
	}
	if _, err := row.Pick(4, caseGrounds); err != nil {
		return caseException{}, err
	}
	return caseException{Graph: row.Cells[0], Word: row.Cells[1], Reading: row.Opt(2), GroundKind: row.Cells[4], Ground: row.Cells[5], Row: row}, nil
}
